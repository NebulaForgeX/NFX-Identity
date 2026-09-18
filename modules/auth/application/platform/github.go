package platform

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"nfxidentity/errors/src/auth"
	"strconv"
	"strings"
	"time"

	"nfxidentity/modules/auth/domain/account"
	"nfxidentity/modules/auth/domain/email"
	"nfxidentity/modules/auth/domain/forgerprofile"
	"nfxidentity/modules/auth/domain/identity"
	"nfxidentity/modules/auth/domain/settings"
	"nfxidentity/pkgs/transaction"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

func (s *Service) oauthCfg() *oauth2.Config {
	if s.github == nil {
		return nil
	}
	return &oauth2.Config{
		ClientID:     s.github.ClientID,
		ClientSecret: s.github.ClientSecret,
		RedirectURL:  s.github.RedirectURL,
		Endpoint:     github.Endpoint,
		Scopes:       []string{"read:user", "user:email"},
	}
}

func (s *Service) GitHubAuthorizeURL(ctx context.Context) (string, string, error) {
	cfg := s.oauthCfg()
	if cfg == nil || cfg.ClientID == "" {
		return "", "", auth.ErrGitHubNotConfigured
	}
	b := make([]byte, 16)
	_, _ = rand.Read(b)
	state := hex.EncodeToString(b)
	if s.redis != nil {
		_ = s.redis.Set(ctx, "nfxidentity:github:state:"+state, "1", 10*time.Minute).Err()
	}
	return cfg.AuthCodeURL(state, oauth2.AccessTypeOnline), state, nil
}

type githubUser struct {
	ID    int64  `json:"id"`
	Login string `json:"login"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (s *Service) exchangeGitHub(ctx context.Context, code string) (*githubUser, error) {
	cfg := s.oauthCfg()
	if cfg == nil || cfg.ClientID == "" {
		return nil, auth.ErrGitHubNotConfigured
	}
	tok, err := cfg.Exchange(ctx, code)
	if err != nil {
		return nil, auth.ErrGitHubExchangeFailed.WithCause(err)
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://api.github.com/user", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+tok.AccessToken)
	req.Header.Set("Accept", "application/vnd.github+json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, auth.ErrGitHubUserFailed.WithCause(err)
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	if resp.StatusCode >= 300 {
		return nil, auth.ErrGitHubUserFailed.WithMsg(fmt.Sprintf("status %d", resp.StatusCode))
	}
	var user githubUser
	if err := json.Unmarshal(body, &user); err != nil {
		return nil, err
	}
	if user.Email == "" {
		user.Email = s.fetchGitHubEmail(ctx, tok.AccessToken)
	}
	return &user, nil
}

func (s *Service) fetchGitHubEmail(ctx context.Context, accessToken string) string {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://api.github.com/user/emails", nil)
	if err != nil {
		return ""
	}
	req.Header.Set("Authorization", "Bearer "+accessToken)
	req.Header.Set("Accept", "application/vnd.github+json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return ""
	}
	defer resp.Body.Close()
	var emails []struct {
		Email    string `json:"email"`
		Primary  bool   `json:"primary"`
		Verified bool   `json:"verified"`
	}
	if json.NewDecoder(resp.Body).Decode(&emails) != nil {
		return ""
	}
	for _, e := range emails {
		if e.Primary && e.Verified {
			return e.Email
		}
	}
	for _, e := range emails {
		if e.Verified {
			return e.Email
		}
	}
	return ""
}

func (s *Service) consumeGitHubState(ctx context.Context, state string) error {
	state = strings.TrimSpace(state)
	if state == "" {
		return auth.ErrInvalidOAuthState
	}
	if s.redis == nil {
		return auth.ErrInvalidOAuthState
	}
	ok, err := s.redis.Get(ctx, "nfxidentity:github:state:"+state).Result()
	if err != nil || ok == "" {
		return auth.ErrInvalidOAuthState
	}
	_ = s.redis.Del(ctx, "nfxidentity:github:state:"+state).Err()
	return nil
}

func (s *Service) LoginWithGitHub(ctx context.Context, code, state, deviceID, platformName string) (*LoginOutput, error) {
	if err := s.consumeGitHubState(ctx, state); err != nil {
		return nil, err
	}
	user, err := s.exchangeGitHub(ctx, code)
	if err != nil {
		return nil, err
	}
	subject := strconv.FormatInt(user.ID, 10)
	ident, err := s.repoFactory.Identity(none()).Get.ByProviderSubject(ctx, "github", subject)
	if err == nil {
		ident.TouchLogin(time.Now())
		_ = s.repoFactory.Identity(none()).Update.Generic(ctx, ident)
		emailAddr, phone := s.primaryContacts(ctx, ident.AccountID())
		if emailAddr == "" {
			emailAddr = user.Email
		}
		return s.issueAccountSession(ctx, ident.AccountID(), ptrUUID(ident.ID()), deviceID, emailAddr, phone)
	}
	if !isMissing(err) {
		return nil, auth.ErrGitHubLookupFailed.WithCause(err)
	}
	if platformName == "" {
		platformName = "nfxidentity"
	}
	now := time.Now()
	accountID := uuid.New()
	identityID := uuid.New()
	profileID := uuid.New()
	display := user.Name
	if display == "" {
		display = user.Login
	}
	err = s.tx.WithUoW(ctx, func(ctx context.Context, uow transaction.UoW) error {
		accountRepo := s.repoFactory.Account(uow)
		identityRepo := s.repoFactory.Identity(uow)
		emailRepo := s.repoFactory.Email(uow)
		forgerRepo := s.repoFactory.Forger(uow)
		settingsRepo := s.repoFactory.Settings(uow)
		if err := accountRepo.Create.New(ctx, account.NewFromState(account.AccountState{
			ID: accountID, AccountStatus: "active", SignupPlatform: platformName, CreatedAt: now, UpdatedAt: now,
		})); err != nil {
			return err
		}
		if err := identityRepo.Create.New(ctx, identity.NewFromState(identity.IdentityState{
			ID: identityID, AccountID: accountID, IdentityProvider: "github", ProviderSubject: subject,
			CreatedAt: now, UpdatedAt: now,
		})); err != nil {
			return err
		}
		if user.Email != "" {
			verified := now
			if err := emailRepo.Create.New(ctx, email.NewFromState(email.EmailState{
				ID: uuid.New(), AccountID: accountID, Address: strings.ToLower(user.Email), IsPrimary: true, VerifiedAt: &verified, CreatedAt: now, UpdatedAt: now,
			})); err != nil {
				return err
			}
		}
		if err := forgerRepo.Create.New(ctx, forgerprofile.NewFromState(forgerprofile.State{
			ID: profileID, AccountID: accountID, Roles: pq.StringArray{"forger"},
			ProfileLanguage: "zh", DisplayName: &display, CreatedAt: now, UpdatedAt: now,
		})); err != nil {
			return err
		}
		return settingsRepo.Create.New(ctx, settings.NewFromState(settings.State{
			ID: profileID, Kind: "forger", LoginNotification: true, CreatedAt: now, UpdatedAt: now,
		}))
	})
	if err != nil {
		return nil, auth.ErrGitHubSignupFailed.WithCause(err)
	}
	return s.issueAccountSession(ctx, accountID, &identityID, deviceID, user.Email, "")
}

func (s *Service) LinkGitHub(ctx context.Context, accountID uuid.UUID, code, state string) error {
	if err := s.consumeGitHubState(ctx, state); err != nil {
		return err
	}
	user, err := s.exchangeGitHub(ctx, code)
	if err != nil {
		return err
	}
	subject := strconv.FormatInt(user.ID, 10)
	existing, err := s.repoFactory.Identity(none()).Get.ByProviderSubject(ctx, "github", subject)
	if err == nil {
		if existing.AccountID() != accountID {
			return auth.ErrGitHubTaken
		}
		return nil
	}
	if !isMissing(err) {
		return err
	}
	now := time.Now()
	return s.repoFactory.Identity(none()).Create.New(ctx, identity.NewFromState(identity.IdentityState{
		ID: uuid.New(), AccountID: accountID, IdentityProvider: "github", ProviderSubject: subject,
		CreatedAt: now, UpdatedAt: now,
	}))
}

func (s *Service) UnlinkGitHub(ctx context.Context, accountID uuid.UUID) error {
	idents, err := s.repoFactory.Identity(none()).Get.ByAccountID(ctx, accountID)
	if err != nil {
		return err
	}
	hasPassword := false
	var githubIdent *identity.Identity
	for _, item := range idents {
		if item.IdentityProvider() == "password" {
			hasPassword = true
		}
		if item.IdentityProvider() == "github" {
			githubIdent = item
		}
	}
	if !hasPassword {
		return auth.ErrLastIdentity
	}
	if githubIdent == nil {
		return nil
	}
	githubIdent.SoftDelete(time.Now())
	return s.repoFactory.Identity(none()).Update.Generic(ctx, githubIdent)
}
