package platform

import (
	"context"
	"encoding/json"
	"time"

	"nfxidentity/pkgs/errx"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func (s *Service) ChangePassword(ctx context.Context, accountID uuid.UUID, current, next, code string) error {
	emailAddr, _ := s.primaryContacts(ctx, accountID)
	if emailAddr == "" || !s.checkVerificationCode(ctx, emailAddr, code) {
		return errx.InvalidArg("INVALID_VERIFICATION_CODE", "invalid verification code")
	}
	idents, err := s.repos.Identity(none()).Get.ByAccountID(ctx, accountID)
	if err != nil {
		return ErrNotFound
	}
	for _, item := range idents {
		if item.IdentityProvider() != "password" {
			continue
		}
		if item.PasswordHash() == nil || bcrypt.CompareHashAndPassword([]byte(*item.PasswordHash()), []byte(current)) != nil {
			return ErrInvalidCredentials
		}
		hash, err := bcrypt.GenerateFromPassword([]byte(next), bcrypt.DefaultCost)
		if err != nil {
			return errx.Internal("HASH_FAILED", "failed to hash password")
		}
		item.SetPasswordHash(string(hash))
		return s.repos.Identity(none()).Update.Generic(ctx, item)
	}
	return ErrNotFound
}

func (s *Service) SendPasswordCode(ctx context.Context, accountID uuid.UUID) error {
	emailAddr, _ := s.primaryContacts(ctx, accountID)
	if emailAddr == "" {
		return ErrNotFound
	}
	s.StoreVerificationCode(ctx, emailAddr, RandomCode())
	return nil
}

func (s *Service) GetAccountByID(ctx context.Context, accountID uuid.UUID) (map[string]any, error) {
	return s.FullAccountWithProfile(ctx, accountID, uuid.Nil, "forger")
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
	acc, err := s.repos.Account(none()).Get.ByID(ctx, accountID)
	if err != nil {
		return nil, ErrNotFound
	}
	emails, _ := s.emails.List.ByAccountID(ctx, accountID)
	phones, _ := s.phones.List.ByAccountID(ctx, accountID)
	idents, _ := s.repos.Identity(none()).Get.ByAccountID(ctx, accountID)
	identJSON := make([]map[string]any, 0, len(idents))
	for _, e := range idents {
		identJSON = append(identJSON, map[string]any{
			"identity_provider": e.IdentityProvider(), "provider_subject": e.ProviderSubject(), "last_login_at": e.LastLoginAt(),
		})
	}
	out := map[string]any{
		"account": map[string]any{
			"id": acc.ID().String(), "account_status": acc.AccountStatus(), "signup_platform": acc.SignupPlatform(),
			"created_at": acc.CreatedAt(), "updated_at": acc.UpdatedAt(),
		},
		"emails":     emailMaps(emails),
		"phones":     phoneMaps(phones),
		"identities": identJSON,
	}
	if kind == "authority" {
		if p, err := s.repos.Authority(none()).Get.ByAccountAndID(ctx, accountID, profileID); err == nil {
			m := s.authorityVO(p)
			m["first_name"] = p.FirstName()
			m["last_name"] = p.LastName()
			m["bio"] = p.Bio()
			m["preference"] = p.Preference()
			out["authority_profile"] = m
		} else {
			out["authority_profile"] = nil
		}
	} else {
		if p, err := s.repos.Forger(none()).Get.ByAccountAndID(ctx, accountID, profileID); err == nil {
			m := s.forgerVO(p)
			m["first_name"] = p.FirstName()
			m["last_name"] = p.LastName()
			m["bio"] = p.Bio()
			m["gender"] = p.Gender()
			m["birthday"] = p.Birthday()
			m["preference"] = p.Preference()
			if avatars, err := s.repos.Avatar(none()).Get.ByProfileID(ctx, "forger", p.ID()); err == nil {
				list := make([]map[string]any, 0, len(avatars))
				for _, a := range avatars {
					list = append(list, map[string]any{"id": a.ID().String(), "image_id": a.ImageID().String(), "is_active": a.IsActive()})
				}
				m["avatars"] = list
			}
			if bgs, err := s.repos.Background(none()).Get.ByProfileID(ctx, "forger", p.ID()); err == nil {
				list := make([]map[string]any, 0, len(bgs))
				for _, b := range bgs {
					list = append(list, map[string]any{"id": b.ID().String(), "image_id": b.ImageID().String(), "sort_order": b.SortOrder()})
				}
				m["backgrounds"] = list
			}
			if st, err := s.repos.Settings(none()).Get.ByID(ctx, "forger", p.ID()); err == nil {
				m["settings"] = map[string]any{"id": st.ID().String(), "login_notification": st.LoginNotification()}
			}
			out["forger_profile"] = m
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
