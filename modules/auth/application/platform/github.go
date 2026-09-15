package platform

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

	"nfxidentity/modules/auth/infrastructure/rdb"
	"nfxidentity/pkgs/errx"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
	"gorm.io/gorm"
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
		return "", "", errx.FailedPrecond("GITHUB_NOT_CONFIGURED", "configure GITHUB_CLIENT_ID to enable GitHub login")
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
		return nil, errx.FailedPrecond("GITHUB_NOT_CONFIGURED", "configure GITHUB_CLIENT_ID to enable GitHub login")
	}
	tok, err := cfg.Exchange(ctx, code)
	if err != nil {
		return nil, errx.Unauthorized("GITHUB_EXCHANGE_FAILED", err.Error())
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://api.github.com/user", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+tok.AccessToken)
	req.Header.Set("Accept", "application/vnd.github+json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, errx.Internal("GITHUB_USER_FAILED", err.Error())
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	if resp.StatusCode >= 300 {
		return nil, errx.Unauthorized("GITHUB_USER_FAILED", fmt.Sprintf("status %d", resp.StatusCode))
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

func (s *Service) LoginWithGitHub(ctx context.Context, code, state, deviceID, platform string) (*LoginOutput, error) {
	if s.redis != nil && state != "" {
		ok, _ := s.redis.Get(ctx, "nfxidentity:github:state:"+state).Result()
		if ok == "" {
			return nil, errx.Unauthorized("INVALID_OAUTH_STATE", "invalid oauth state")
		}
		_ = s.redis.Del(ctx, "nfxidentity:github:state:"+state).Err()
	}
	user, err := s.exchangeGitHub(ctx, code)
	if err != nil {
		return nil, err
	}
	subject := strconv.FormatInt(user.ID, 10)
	var ident rdb.Identity
	err = s.db.WithContext(ctx).Where("identity_provider = ? AND provider_subject = ? AND deleted_at IS NULL", "github", subject).First(&ident).Error
	if err == nil {
		now := time.Now()
		_ = s.db.WithContext(ctx).Model(&ident).Update("last_login_at", now).Error
		email, phone := s.primaryContacts(ctx, ident.AccountID)
		if email == "" {
			email = user.Email
		}
		return s.issueAccountSession(ctx, ident.AccountID, &ident.ID, deviceID, email, phone)
	}
	if err != gorm.ErrRecordNotFound {
		return nil, errx.Internal("GITHUB_LOOKUP_FAILED", err.Error())
	}
	if platform == "" {
		platform = "nfxidentity"
	}
	now := time.Now()
	accountID := uuid.New()
	identityID := uuid.New()
	profileID := uuid.New()
	display := user.Name
	if display == "" {
		display = user.Login
	}
	err = s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&rdb.Account{
			ID: accountID, AccountStatus: "active", SignupPlatform: platform, CreatedAt: now, UpdatedAt: now,
		}).Error; err != nil {
			return err
		}
		if err := tx.Create(&rdb.Identity{
			ID: identityID, AccountID: accountID, IdentityProvider: "github", ProviderSubject: subject,
			CreatedAt: now, UpdatedAt: now,
		}).Error; err != nil {
			return err
		}
		if user.Email != "" {
			verified := now
			if err := tx.Create(&rdb.Email{
				ID: uuid.New(), AccountID: accountID, Email: strings.ToLower(user.Email), IsPrimary: true, VerifiedAt: &verified, CreatedAt: now, UpdatedAt: now,
			}).Error; err != nil {
				return err
			}
		}
		if err := tx.Create(&rdb.ForgerProfile{
			ID: profileID, AccountID: accountID, ForgerRoles: pq.StringArray{"forger"},
			ProfileLanguage: "zh", DisplayName: &display, CreatedAt: now, UpdatedAt: now,
		}).Error; err != nil {
			return err
		}
		return tx.Create(&rdb.ForgerProfileSettings{ProfileSettings: rdb.ProfileSettings{
			ID: profileID, LoginNotification: true, CreatedAt: now, UpdatedAt: now,
		}}).Error
	})
	if err != nil {
		return nil, errx.Internal("GITHUB_SIGNUP_FAILED", err.Error())
	}
	return s.issueAccountSession(ctx, accountID, &identityID, deviceID, user.Email, "")
}

func (s *Service) LinkGitHub(ctx context.Context, accountID uuid.UUID, code string) error {
	user, err := s.exchangeGitHub(ctx, code)
	if err != nil {
		return err
	}
	subject := strconv.FormatInt(user.ID, 10)
	var existing rdb.Identity
	if err := s.db.WithContext(ctx).Where("identity_provider = ? AND provider_subject = ? AND deleted_at IS NULL", "github", subject).First(&existing).Error; err == nil {
		if existing.AccountID != accountID {
			return errx.Conflict("GITHUB_TAKEN", "github account already linked")
		}
		return nil
	}
	now := time.Now()
	return s.db.WithContext(ctx).Create(&rdb.Identity{
		ID: uuid.New(), AccountID: accountID, IdentityProvider: "github", ProviderSubject: subject,
		CreatedAt: now, UpdatedAt: now,
	}).Error
}

func (s *Service) UnlinkGitHub(ctx context.Context, accountID uuid.UUID) error {
	var n int64
	s.db.WithContext(ctx).Model(&rdb.Identity{}).Where("account_id = ? AND identity_provider = ? AND deleted_at IS NULL", accountID, "password").Count(&n)
	if n == 0 {
		return errx.FailedPrecond("LAST_IDENTITY", "cannot unlink the last identity")
	}
	now := time.Now()
	return s.db.WithContext(ctx).Model(&rdb.Identity{}).
		Where("account_id = ? AND identity_provider = ? AND deleted_at IS NULL", accountID, "github").
		Update("deleted_at", now).Error
}
