package platform

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"nfxidentity/modules/auth/infrastructure/rdb"
	"nfxidentity/pkgs/errx"
	"nfxidentity/pkgs/tokenx"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"github.com/redis/go-redis/v9"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var (
	ErrInvalidCredentials = errx.Unauthorized("INVALID_CREDENTIALS", "invalid credentials")
	ErrInvalidToken       = errx.Unauthorized("INVALID_TOKEN", "invalid or expired token")
	ErrInvalidRefresh     = errx.Unauthorized("INVALID_REFRESH_TOKEN", "invalid refresh token")
	ErrProfileNotOwned    = errx.Forbidden("PROFILE_NOT_OWNED", "profile is not owned by account")
	ErrOwnerRequired      = errx.Forbidden("AUTHORITY_PROFILE_INSUFFICIENT_ROLE", "owner role required")
	ErrNotFound           = errx.NotFound("NOT_FOUND", "resource not found")
)

const emptyProfileID = "00000000-0000-0000-0000-000000000000"

type Service struct {
	db      *gorm.DB
	tokens  *tokenx.Tokenx
	github  *GitHubConfig
	redis   *redis.Client
	checkFn checkFunc
}

type GitHubConfig struct {
	ClientID     string
	ClientSecret string
	RedirectURL  string
}

func NewService(db *gorm.DB, tokens *tokenx.Tokenx, github *GitHubConfig, redisClient *redis.Client) *Service {
	s := &Service{db: db, tokens: tokens, github: github, redis: redisClient}
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

func (s *Service) SignupWithEmail(ctx context.Context, email, password, code, lang, deviceID, platform string) (*LoginOutput, error) {
	email = strings.ToLower(strings.TrimSpace(email))
	if email == "" || password == "" {
		return nil, errx.InvalidArg("INVALID_PARAMS", "email and password required")
	}
	if !s.checkVerificationCode(ctx, email, code) {
		return nil, errx.InvalidArg("INVALID_VERIFICATION_CODE", "invalid verification code")
	}
	if platform == "" {
		platform = "nfxidentity"
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errx.Internal("HASH_FAILED", "failed to hash password")
	}
	now := time.Now()
	accountID := uuid.New()
	identityID := uuid.New()
	emailID := uuid.New()
	profileID := uuid.New()
	hashStr := string(hash)
	verified := now

	err = s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&rdb.Account{
			ID: accountID, AccountStatus: "active", SignupPlatform: platform, CreatedAt: now, UpdatedAt: now,
		}).Error; err != nil {
			return err
		}
		if err := tx.Create(&rdb.Identity{
			ID: identityID, AccountID: accountID, IdentityProvider: "password", ProviderSubject: email,
			PasswordHash: &hashStr, CreatedAt: now, UpdatedAt: now,
		}).Error; err != nil {
			return err
		}
		if err := tx.Create(&rdb.Email{
			ID: emailID, AccountID: accountID, Email: email, IsPrimary: true, VerifiedAt: &verified, CreatedAt: now, UpdatedAt: now,
		}).Error; err != nil {
			return err
		}
		display := email
		if err := tx.Create(&rdb.ForgerProfile{
			ID: profileID, AccountID: accountID, ForgerRoles: pq.StringArray{"forger"},
			ProfileLanguage: langOrDefault(lang), DisplayName: &display, CreatedAt: now, UpdatedAt: now,
		}).Error; err != nil {
			return err
		}
		return tx.Create(&rdb.ForgerProfileSettings{ProfileSettings: rdb.ProfileSettings{
			ID: profileID, LoginNotification: true, CreatedAt: now, UpdatedAt: now,
		}}).Error
	})
	if err != nil {
		if strings.Contains(err.Error(), "uq_emails") {
			return nil, errx.Conflict("EMAIL_TAKEN", "email already registered")
		}
		return nil, errx.Internal("SIGNUP_FAILED", err.Error())
	}
	return s.issueAccountSession(ctx, accountID, &identityID, deviceID, email, "")
}

func (s *Service) LoginWithEmail(ctx context.Context, email, password, deviceID string) (*LoginOutput, error) {
	email = strings.ToLower(strings.TrimSpace(email))
	var ident rdb.Identity
	err := s.db.WithContext(ctx).Where("identity_provider = ? AND provider_subject = ? AND deleted_at IS NULL", "password", email).First(&ident).Error
	if err != nil || ident.PasswordHash == nil {
		return nil, ErrInvalidCredentials
	}
	if bcrypt.CompareHashAndPassword([]byte(*ident.PasswordHash), []byte(password)) != nil {
		return nil, ErrInvalidCredentials
	}
	now := time.Now()
	_ = s.db.WithContext(ctx).Model(&ident).Update("last_login_at", now).Error
	return s.issueAccountSession(ctx, ident.AccountID, &ident.ID, deviceID, email, "")
}

func (s *Service) LoginWithPhone(ctx context.Context, phone, password, deviceID string) (*LoginOutput, error) {
	phone = strings.TrimSpace(phone)
	var p rdb.Phone
	if err := s.db.WithContext(ctx).Where("phone = ? AND deleted_at IS NULL", phone).First(&p).Error; err != nil {
		return nil, ErrInvalidCredentials
	}
	var ident rdb.Identity
	if err := s.db.WithContext(ctx).Where("account_id = ? AND identity_provider = ? AND deleted_at IS NULL", p.AccountID, "password").First(&ident).Error; err != nil {
		return nil, ErrInvalidCredentials
	}
	if ident.PasswordHash == nil || bcrypt.CompareHashAndPassword([]byte(*ident.PasswordHash), []byte(password)) != nil {
		return nil, ErrInvalidCredentials
	}
	return s.issueAccountSession(ctx, p.AccountID, &ident.ID, deviceID, "", phone)
}

func (s *Service) SelectProfile(ctx context.Context, accountID uuid.UUID, profileID uuid.UUID, kind, deviceID string) (*SelectProfileOutput, error) {
	email, phone := s.primaryContacts(ctx, accountID)
	display := ""
	switch kind {
	case "forger":
		var p rdb.ForgerProfile
		if err := s.db.WithContext(ctx).Where("id = ? AND account_id = ? AND deleted_at IS NULL", profileID, accountID).First(&p).Error; err != nil {
			return nil, ErrProfileNotOwned
		}
		if p.DisplayName != nil {
			display = *p.DisplayName
		}
	case "authority":
		var p rdb.AuthorityProfile
		if err := s.db.WithContext(ctx).Where("id = ? AND account_id = ? AND deleted_at IS NULL", profileID, accountID).First(&p).Error; err != nil {
			return nil, ErrProfileNotOwned
		}
		if p.DisplayName != nil {
			display = *p.DisplayName
		}
	default:
		return nil, errx.InvalidArg("INVALID_PROFILE_KIND", "kind must be forger or authority")
	}
	access, refresh, err := s.tokens.GenerateTokenPair(accountID.String(), profileID.String(), display, email, phone, kind)
	if err != nil {
		return nil, errx.Internal("TOKEN_FAILED", err.Error())
	}
	if err := s.persistRefresh(ctx, accountID, nil, &profileID, &kind, deviceID, refresh); err != nil {
		return nil, err
	}
	return &SelectProfileOutput{
		AccountID: accountID.String(), ProfileID: profileID.String(), AccessToken: access, RefreshToken: refresh,
	}, nil
}

func (s *Service) Refresh(ctx context.Context, refreshToken, deviceID string) (*TokenOutput, error) {
	claims, err := s.tokens.VerifyRefreshToken(refreshToken)
	if err != nil {
		return nil, ErrInvalidRefresh
	}
	hash := hashToken(refreshToken)
	var row rdb.RefreshToken
	if err := s.db.WithContext(ctx).Where("token_hash = ? AND revoked_at IS NULL AND deleted_at IS NULL", hash).First(&row).Error; err != nil {
		return nil, ErrInvalidRefresh
	}
	now := time.Now()
	_ = s.db.WithContext(ctx).Model(&row).Update("revoked_at", now).Error
	access, refresh, err := s.tokens.GenerateTokenPair(claims.AccountID, claims.ProfileID, claims.Username, claims.Email, claims.Phone, claims.ProfileScope)
	if err != nil {
		return nil, errx.Internal("TOKEN_FAILED", err.Error())
	}
	accountID, _ := uuid.Parse(claims.AccountID)
	var profileID *uuid.UUID
	if claims.ProfileID != "" && claims.ProfileID != emptyProfileID {
		id, err := uuid.Parse(claims.ProfileID)
		if err == nil {
			profileID = &id
		}
	}
	scope := claims.ProfileScope
	if err := s.persistRefresh(ctx, accountID, row.IdentityID, profileID, nullableStr(scope), deviceID, refresh); err != nil {
		return nil, err
	}
	return &TokenOutput{AccessToken: access, RefreshToken: refresh}, nil
}

func (s *Service) ListOwnerForgerProfiles(ctx context.Context, accountID uuid.UUID, query string, limit, offset int) (items []map[string]any, total int64, err error) {
	if err := s.requireOwner(ctx, accountID); err != nil {
		return nil, 0, err
	}
	q := s.db.WithContext(ctx).Model(&rdb.ForgerProfile{}).Where("deleted_at IS NULL")
	if query != "" {
		like := "%" + query + "%"
		q = q.Where("display_name ILIKE ? OR city ILIKE ?", like, like)
	}
	if err := q.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	var rows []rdb.ForgerProfile
	if err := q.Order("created_at desc").Limit(limitOr(limit, 20)).Offset(offset).Find(&rows).Error; err != nil {
		return nil, 0, err
	}
	for _, r := range rows {
		items = append(items, s.forgerItem(r))
	}
	return items, total, nil
}

func (s *Service) ListOwnerAuthorityProfiles(ctx context.Context, accountID uuid.UUID, query string, limit, offset int) (items []map[string]any, total int64, err error) {
	if err := s.requireOwner(ctx, accountID); err != nil {
		return nil, 0, err
	}
	q := s.db.WithContext(ctx).Model(&rdb.AuthorityProfile{}).Where("deleted_at IS NULL")
	if query != "" {
		like := "%" + query + "%"
		q = q.Where("display_name ILIKE ?", like)
	}
	if err := q.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	var rows []rdb.AuthorityProfile
	if err := q.Order("created_at desc").Limit(limitOr(limit, 20)).Offset(offset).Find(&rows).Error; err != nil {
		return nil, 0, err
	}
	for _, r := range rows {
		items = append(items, s.authorityItem(r))
	}
	return items, total, nil
}

func (s *Service) EnsureOwnedProfile(ctx context.Context, accountID, profileID, scope string) error {
	aid, err := uuid.Parse(accountID)
	if err != nil {
		return ErrProfileNotOwned
	}
	pid, err := uuid.Parse(profileID)
	if err != nil {
		return ErrProfileNotOwned
	}
	switch scope {
	case "forger":
		var n int64
		s.db.WithContext(ctx).Model(&rdb.ForgerProfile{}).Where("id = ? AND account_id = ? AND deleted_at IS NULL", pid, aid).Count(&n)
		if n == 0 {
			return ErrProfileNotOwned
		}
	case "authority":
		var n int64
		s.db.WithContext(ctx).Model(&rdb.AuthorityProfile{}).Where("id = ? AND account_id = ? AND deleted_at IS NULL", pid, aid).Count(&n)
		if n == 0 {
			return ErrProfileNotOwned
		}
	default:
		return ErrProfileNotOwned
	}
	return nil
}

func (s *Service) HasAuthorityRole(ctx context.Context, profileID, role string) (bool, error) {
	pid, err := uuid.Parse(profileID)
	if err != nil {
		return false, err
	}
	var p rdb.AuthorityProfile
	if err := s.db.WithContext(ctx).Where("id = ? AND deleted_at IS NULL", pid).First(&p).Error; err != nil {
		return false, nil
	}
	for _, r := range p.AuthorityRoles {
		if r == role {
			return true, nil
		}
	}
	return false, nil
}

func (s *Service) HasForgerRole(ctx context.Context, profileID, role string) (bool, error) {
	pid, err := uuid.Parse(profileID)
	if err != nil {
		return false, err
	}
	var p rdb.ForgerProfile
	if err := s.db.WithContext(ctx).Where("id = ? AND deleted_at IS NULL", pid).First(&p).Error; err != nil {
		return false, nil
	}
	for _, r := range p.ForgerRoles {
		if r == role {
			return true, nil
		}
	}
	return false, nil
}

func (s *Service) FullAccountWithProfile(ctx context.Context, accountID uuid.UUID, profileID uuid.UUID, kind string) (map[string]any, error) {
	if s.redis != nil {
		if raw, err := s.redis.Get(ctx, fullCacheKey(accountID, profileID)).Bytes(); err == nil {
			var cached map[string]any
			if json.Unmarshal(raw, &cached) == nil {
				return cached, nil
			}
		}
	}
	var acc rdb.Account
	if err := s.db.WithContext(ctx).Where("id = ? AND deleted_at IS NULL", accountID).First(&acc).Error; err != nil {
		return nil, ErrNotFound
	}
	var emails []rdb.Email
	_ = s.db.WithContext(ctx).Where("account_id = ? AND deleted_at IS NULL", accountID).Find(&emails).Error
	var phones []rdb.Phone
	_ = s.db.WithContext(ctx).Where("account_id = ? AND deleted_at IS NULL", accountID).Find(&phones).Error
	var idents []rdb.Identity
	_ = s.db.WithContext(ctx).Where("account_id = ? AND deleted_at IS NULL", accountID).Find(&idents).Error
	out := map[string]any{
		"account": map[string]any{
			"id": acc.ID.String(), "account_status": acc.AccountStatus, "signup_platform": acc.SignupPlatform,
			"created_at": acc.CreatedAt, "updated_at": acc.UpdatedAt,
		},
		"emails":     emailsJSON(emails),
		"phones":     phonesJSON(phones),
		"identities": identitiesJSON(idents),
	}
	if kind == "authority" {
		var p rdb.AuthorityProfile
		if err := s.db.WithContext(ctx).Where("id = ? AND account_id = ? AND deleted_at IS NULL", profileID, accountID).First(&p).Error; err == nil {
			out["authority_profile"] = s.authorityFull(ctx, p)
		} else {
			out["authority_profile"] = nil
		}
	} else {
		var p rdb.ForgerProfile
		if err := s.db.WithContext(ctx).Where("id = ? AND account_id = ? AND deleted_at IS NULL", profileID, accountID).First(&p).Error; err == nil {
			out["forger_profile"] = s.forgerFull(ctx, p)
		} else {
			out["forger_profile"] = nil
		}
	}
	if s.redis != nil {
		if raw, err := json.Marshal(out); err == nil {
			_ = s.redis.Set(ctx, fullCacheKey(accountID, profileID), raw, 5*time.Minute).Err()
		}
	}
	return out, nil
}

func (s *Service) StoreVerificationCode(ctx context.Context, email, code string) {
	if s.redis == nil {
		return
	}
	_ = s.redis.Set(ctx, verifyKey(email), code, 10*time.Minute).Err()
}

func (s *Service) redisCheckCode(ctx context.Context, email, code string) bool {
	if s.redis == nil {
		return len(code) >= 4
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

var defaultCheck = func(ctx context.Context, email, code string) bool {
	return len(code) >= 4
}

func (s *Service) checkVerificationCode(ctx context.Context, email, code string) bool {
	if s.checkFn != nil {
		return s.checkFn(ctx, email, code)
	}
	return defaultCheck(ctx, email, code)
}

type checkFunc func(ctx context.Context, email, code string) bool

// checkFn is set optionally.
func init() {}

func (s *Service) issueAccountSession(ctx context.Context, accountID uuid.UUID, identityID *uuid.UUID, deviceID, email, phone string) (*LoginOutput, error) {
	access, refresh, err := s.tokens.GenerateTokenPair(accountID.String(), emptyProfileID, "", email, phone, "")
	if err != nil {
		return nil, errx.Internal("TOKEN_FAILED", err.Error())
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
	if deviceID != "" {
		_ = s.db.WithContext(ctx).Model(&rdb.RefreshToken{}).
			Where("account_id = ? AND device_id = ? AND revoked_at IS NULL", accountID, deviceID).
			Update("revoked_at", now).Error
	}
	row := rdb.RefreshToken{
		ID: uuid.New(), AccountID: accountID, IdentityID: identityID, ProfileID: profileID, ProfileScope: scope,
		TokenHash: hashToken(refresh), ExpiresAt: now.Add(30 * 24 * time.Hour), CreatedAt: now,
	}
	if deviceID != "" {
		row.DeviceID = &deviceID
	}
	return s.db.WithContext(ctx).Create(&row).Error
}

func (s *Service) ListProfileItemsExport(ctx context.Context, accountID uuid.UUID) []ProfileItem {
	return s.listProfileItems(ctx, accountID)
}

func (s *Service) listProfileItems(ctx context.Context, accountID uuid.UUID) []ProfileItem {
	var out []ProfileItem
	var forgers []rdb.ForgerProfile
	_ = s.db.WithContext(ctx).Where("account_id = ? AND deleted_at IS NULL", accountID).Find(&forgers).Error
	for _, p := range forgers {
		item := ProfileItem{ProfileID: p.ID.String(), Kind: "forger", Roles: []string(p.ForgerRoles), DisplayName: p.DisplayName, City: p.City, Country: p.Country}
		out = append(out, item)
	}
	var auths []rdb.AuthorityProfile
	_ = s.db.WithContext(ctx).Where("account_id = ? AND deleted_at IS NULL", accountID).Find(&auths).Error
	for _, p := range auths {
		item := ProfileItem{ProfileID: p.ID.String(), Kind: "authority", Roles: []string(p.AuthorityRoles), DisplayName: p.DisplayName, City: p.City, Country: p.Country}
		out = append(out, item)
	}
	return out
}

func (s *Service) PrimaryContactsExport(ctx context.Context, accountID uuid.UUID) (email, phone string) {
	return s.primaryContacts(ctx, accountID)
}

func (s *Service) primaryContacts(ctx context.Context, accountID uuid.UUID) (email, phone string) {
	var e rdb.Email
	if err := s.db.WithContext(ctx).Where("account_id = ? AND is_primary = true AND deleted_at IS NULL", accountID).First(&e).Error; err == nil {
		email = e.Email
	}
	var p rdb.Phone
	if err := s.db.WithContext(ctx).Where("account_id = ? AND is_primary = true AND deleted_at IS NULL", accountID).First(&p).Error; err == nil {
		phone = p.Phone
	}
	return
}

func (s *Service) requireOwner(ctx context.Context, accountID uuid.UUID) error {
	var n int64
	s.db.WithContext(ctx).Model(&rdb.AuthorityProfile{}).
		Where("account_id = ? AND deleted_at IS NULL AND authority_roles @> ARRAY['owner']::auth.authority_role[]", accountID).
		Count(&n)
	if n == 0 {
		return ErrOwnerRequired
	}
	return nil
}

func (s *Service) forgerItem(p rdb.ForgerProfile) map[string]any {
	return map[string]any{
		"profile_id": p.ID.String(), "account_id": p.AccountID.String(), "display_name": p.DisplayName,
		"profile_language": p.ProfileLanguage, "city": p.City, "country": p.Country, "website": p.Website,
		"timezone": p.Timezone, "forger_roles": []string(p.ForgerRoles), "created_at": p.CreatedAt,
	}
}

func (s *Service) authorityItem(p rdb.AuthorityProfile) map[string]any {
	return map[string]any{
		"profile_id": p.ID.String(), "account_id": p.AccountID.String(), "display_name": p.DisplayName,
		"profile_language": p.ProfileLanguage, "city": p.City, "country": p.Country,
		"authority_roles": []string(p.AuthorityRoles), "created_at": p.CreatedAt,
	}
}

func (s *Service) forgerFull(ctx context.Context, p rdb.ForgerProfile) map[string]any {
	m := s.forgerItem(p)
	m["first_name"] = p.FirstName
	m["last_name"] = p.LastName
	m["bio"] = p.Bio
	m["gender"] = p.Gender
	m["birthday"] = p.Birthday
	m["preference"] = json.RawMessage(p.Preference)
	var avatars []rdb.ForgerProfileAvatar
	_ = s.db.WithContext(ctx).Where("profile_id = ? AND deleted_at IS NULL", p.ID).Find(&avatars).Error
	m["avatars"] = avatars
	var bgs []rdb.ForgerProfileBackground
	_ = s.db.WithContext(ctx).Where("profile_id = ? AND deleted_at IS NULL", p.ID).Find(&bgs).Error
	m["backgrounds"] = bgs
	var settings rdb.ForgerProfileSettings
	if err := s.db.WithContext(ctx).Where("id = ?", p.ID).First(&settings).Error; err == nil {
		m["settings"] = settings
	}
	return m
}

func (s *Service) authorityFull(ctx context.Context, p rdb.AuthorityProfile) map[string]any {
	m := s.authorityItem(p)
	m["first_name"] = p.FirstName
	m["last_name"] = p.LastName
	m["bio"] = p.Bio
	m["preference"] = json.RawMessage(p.Preference)
	return m
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

func emailsJSON(rows []rdb.Email) []map[string]any {
	out := make([]map[string]any, 0, len(rows))
	for _, e := range rows {
		out = append(out, map[string]any{
			"id": e.ID.String(), "account_id": e.AccountID.String(), "email": e.Email,
			"is_primary": e.IsPrimary, "verified_at": e.VerifiedAt, "created_at": e.CreatedAt, "updated_at": e.UpdatedAt,
		})
	}
	return out
}

func phonesJSON(rows []rdb.Phone) []map[string]any {
	out := make([]map[string]any, 0, len(rows))
	for _, e := range rows {
		out = append(out, map[string]any{
			"id": e.ID.String(), "account_id": e.AccountID.String(), "phone": e.Phone,
			"is_primary": e.IsPrimary, "verified_at": e.VerifiedAt, "created_at": e.CreatedAt, "updated_at": e.UpdatedAt,
		})
	}
	return out
}

func identitiesJSON(rows []rdb.Identity) []map[string]any {
	out := make([]map[string]any, 0, len(rows))
	for _, e := range rows {
		out = append(out, map[string]any{
			"identity_provider": e.IdentityProvider, "provider_subject": e.ProviderSubject, "last_login_at": e.LastLoginAt,
		})
	}
	return out
}

func RandomCode() string {
	b := make([]byte, 3)
	_, _ = rand.Read(b)
	return fmt.Sprintf("%06d", int(b[0])<<16|int(b[1])<<8|int(b[2])%1000000%1000000)
}
