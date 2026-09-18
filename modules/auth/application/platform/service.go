package platform

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"
	"time"

	"nfxidentity/errors/src/auth"
	repofactory "nfxidentity/modules/auth/infrastructure/repository/factory"
	emailQuery "nfxidentity/modules/auth/query/email"
	phoneQuery "nfxidentity/modules/auth/query/phone"
	profileQuery "nfxidentity/modules/auth/query/profile"
	"nfxidentity/pkgs/email"
	"nfxidentity/pkgs/errx"
	"nfxidentity/pkgs/tokenx"
	"nfxidentity/pkgs/transaction"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

const emptyProfileID = "00000000-0000-0000-0000-000000000000"

type Service struct {
	tx          transaction.TxManager
	repoFactory *repofactory.TxRepoFactory
	emails      *emailQuery.Query
	phones      *phoneQuery.Query
	profiles    *profileQuery.Query
	tokens      *tokenx.Tokenx
	github      *GitHubConfig
	redis       *redis.Client
	mail        *email.EmailService
	checkFn     checkFunc
}

type GitHubConfig struct {
	ClientID     string
	ClientSecret string
	RedirectURL  string
}

func NewService(
	tx transaction.TxManager,
	repoFactory *repofactory.TxRepoFactory,
	emails *emailQuery.Query,
	phones *phoneQuery.Query,
	profiles *profileQuery.Query,
	tokens *tokenx.Tokenx,
	github *GitHubConfig,
	redisClient *redis.Client,
	mail *email.EmailService,
) *Service {
	s := &Service{
		tx: tx, repoFactory: repoFactory, emails: emails, phones: phones, profiles: profiles,
		tokens: tokens, github: github, redis: redisClient, mail: mail,
	}
	if redisClient != nil {
		s.SetVerificationChecker(s.redisCheckCode)
	}
	return s
}

type ProfileItem struct {
	ProfileID     string   `json:"profile_id"`
	Kind          string   `json:"kind"`
	Roles         []string `json:"roles"`
	DisplayName   *string  `json:"display_name"`
	AvatarImageID *string  `json:"avatar_image_id"`
	City          *string  `json:"city"`
	Country       *string  `json:"country"`
}

type LoginOutput struct {
	AccountID    string        `json:"account_id"`
	AccessToken  string        `json:"access_token"`
	RefreshToken string        `json:"refresh_token"`
	Profiles     []ProfileItem `json:"profiles"`
}

type TokenOutput struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type SelectProfileOutput struct {
	AccountID    string `json:"account_id"`
	ProfileID    string `json:"profile_id"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type checkFunc func(ctx context.Context, email, code string) bool

func isMissing(err error) bool {
	e := errx.AsError(err)
	return e != nil && e.Kind == errx.KindNotFound
}

func none() transaction.UoW { return transaction.UoW{} }

func (s *Service) StoreVerificationCode(ctx context.Context, email, code string) error {
	if s.redis == nil {
		return auth.ErrVerificationCodeSaveFailed
	}
	if err := s.redis.Set(ctx, verifyKey(email), code, 10*time.Minute).Err(); err != nil {
		return auth.ErrVerificationCodeSaveFailed.WithCause(err)
	}
	return nil
}

func (s *Service) redisCheckCode(ctx context.Context, email, code string) bool {
	if s.redis == nil {
		return false
	}
	got, err := s.redis.Get(ctx, verifyKey(email)).Result()
	if err != nil {
		return false
	}
	ok := got == code
	if ok {
		_ = s.redis.Del(ctx, verifyKey(email)).Err()
	}
	return ok
}

func verifyKey(email string) string {
	return "nfxidentity:verify:" + strings.ToLower(strings.TrimSpace(email))
}

func (s *Service) SetVerificationChecker(fn func(ctx context.Context, email, code string) bool) {
	s.checkFn = fn
}

func (s *Service) consumeVerificationCode(ctx context.Context, email, code string) error {
	code = strings.TrimSpace(code)
	if code == "" {
		return auth.ErrVerificationCodeWrong
	}
	if s.checkFn != nil {
		if s.checkFn(ctx, email, code) {
			return nil
		}
		return auth.ErrVerificationCodeWrong
	}
	return auth.ErrVerificationCodeExpired
}

func (s *Service) issueAccountSession(ctx context.Context, accountID uuid.UUID, identityID *uuid.UUID, deviceID, email, phone string) (*LoginOutput, error) {
	access, refresh, err := s.tokens.GenerateTokenPair(accountID.String(), emptyProfileID, "", email, phone, "")
	if err != nil {
		return nil, auth.ErrTokenFailed.WithCause(err)
	}
	if err := s.persistRefresh(ctx, accountID, identityID, nil, nil, deviceID, refresh); err != nil {
		return nil, err
	}
	profiles := s.listProfileItems(ctx, accountID)
	return &LoginOutput{
		AccountID: accountID.String(), AccessToken: access, RefreshToken: refresh, Profiles: profiles,
	}, nil
}

func (s *Service) persistRefresh(ctx context.Context, accountID uuid.UUID, identityID, profileID *uuid.UUID, scope *string, deviceID, refresh string) error {
	now := time.Now()
	return s.tx.WithUoW(ctx, func(ctx context.Context, uow transaction.UoW) error {
		refreshTokenRepo := s.repoFactory.RefreshToken(uow)
		if deviceID != "" {
			if err := refreshTokenRepo.Update.RevokeDevice(ctx, accountID, deviceID, now); err != nil {
				return err
			}
		}
		var dev *string
		if deviceID != "" {
			dev = &deviceID
		}
		row := refreshtokenNew(accountID, identityID, profileID, scope, dev, hashToken(refresh), now)
		return refreshTokenRepo.Create.New(ctx, row)
	})
}

func (s *Service) ListProfileItemsExport(ctx context.Context, accountID uuid.UUID) []ProfileItem {
	return s.listProfileItems(ctx, accountID)
}

func (s *Service) listProfileItems(ctx context.Context, accountID uuid.UUID) []ProfileItem {
	var out []ProfileItem
	forgers, _ := s.profiles.Forger.ByAccountID(ctx, accountID)
	for _, p := range forgers {
		item := ProfileItem{ProfileID: p.ProfileID.String(), Kind: "forger", Roles: []string(p.ForgerRoles), DisplayName: p.DisplayName, City: p.City, Country: p.Country}
		if p.AvatarImageID != nil {
			id := p.AvatarImageID.String()
			item.AvatarImageID = &id
		}
		out = append(out, item)
	}
	auths, _ := s.profiles.Authority.ByAccountID(ctx, accountID)
	for _, p := range auths {
		item := ProfileItem{ProfileID: p.ProfileID.String(), Kind: "authority", Roles: []string(p.AuthorityRoles), DisplayName: p.DisplayName, City: p.City, Country: p.Country}
		if p.AvatarImageID != nil {
			id := p.AvatarImageID.String()
			item.AvatarImageID = &id
		}
		out = append(out, item)
	}
	return out
}

func (s *Service) PrimaryContactsExport(ctx context.Context, accountID uuid.UUID) (email, phone string) {
	return s.primaryContacts(ctx, accountID)
}

func (s *Service) primaryContacts(ctx context.Context, accountID uuid.UUID) (email, phone string) {
	rows, _ := s.emails.List.ByAccountID(ctx, accountID)
	for _, e := range rows {
		if e.IsPrimary {
			email = e.Email
			break
		}
	}
	phones, _ := s.phones.List.ByAccountID(ctx, accountID)
	for _, p := range phones {
		if p.IsPrimary {
			phone = p.Phone
			break
		}
	}
	return
}

func (s *Service) requireOwner(ctx context.Context, accountID uuid.UUID) error {
	ok, err := s.repoFactory.Authority(none()).Get.HasOwnerRole(ctx, accountID)
	if err != nil {
		return err
	}
	if !ok {
		return auth.ErrAuthorityProfileInsufficientRole
	}
	return nil
}

func (s *Service) EnsureOwnedProfile(ctx context.Context, accountID, profileID, scope string) error {
	aid, err := uuid.Parse(accountID)
	if err != nil {
		return auth.ErrProfileNotOwned
	}
	pid, err := uuid.Parse(profileID)
	if err != nil {
		return auth.ErrProfileNotOwned
	}
	switch scope {
	case "forger":
		ok, err := s.repoFactory.Forger(none()).Get.Owned(ctx, aid, pid)
		if err != nil || !ok {
			return auth.ErrProfileNotOwned
		}
	case "authority":
		ok, err := s.repoFactory.Authority(none()).Get.Owned(ctx, aid, pid)
		if err != nil || !ok {
			return auth.ErrProfileNotOwned
		}
	default:
		return auth.ErrProfileNotOwned
	}
	return nil
}

func hashToken(token string) string {
	sum := sha256.Sum256([]byte(token))
	return hex.EncodeToString(sum[:])
}

func langOrDefault(lang string) string {
	switch lang {
	case "en", "zh", "fr":
		return lang
	default:
		return "zh"
	}
}

func limitOr(n, d int) int {
	if n <= 0 {
		return d
	}
	return n
}

func nullableStr(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}

func RandomCode() string {
	b := make([]byte, 3)
	_, _ = rand.Read(b)
	n := (int(b[0])<<16 | int(b[1])<<8 | int(b[2])) % 1000000
	return fmt.Sprintf("%06d", n)
}

func fullCacheKey(accountID, profileID uuid.UUID) string {
	return "nfxidentity:full:" + accountID.String() + ":" + profileID.String()
}

func (s *Service) InvalidateFull(ctx context.Context, accountID, profileID uuid.UUID) error {
	if s.redis == nil {
		return nil
	}
	return s.redis.Del(ctx, fullCacheKey(accountID, profileID)).Err()
}

func emailMaps(rows []emailQuery.EmailItemVO) []map[string]any {
	out := make([]map[string]any, 0, len(rows))
	for _, e := range rows {
		out = append(out, map[string]any{
			"id": e.ID.String(), "account_id": e.AccountID.String(), "email": e.Email,
			"is_primary": e.IsPrimary, "verified_at": e.VerifiedAt, "created_at": e.CreatedAt, "updated_at": e.UpdatedAt,
		})
	}
	return out
}

func phoneMaps(rows []phoneQuery.PhoneItemVO) []map[string]any {
	out := make([]map[string]any, 0, len(rows))
	for _, e := range rows {
		out = append(out, map[string]any{
			"id": e.ID.String(), "account_id": e.AccountID.String(), "phone": e.Phone,
			"is_primary": e.IsPrimary, "verified_at": e.VerifiedAt, "created_at": e.CreatedAt, "updated_at": e.UpdatedAt,
		})
	}
	return out
}
