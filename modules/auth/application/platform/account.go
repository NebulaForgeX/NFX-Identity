package platform

import (
	"context"
	"nfxidentity/errors/src/auth"
	"nfxidentity/errors/src/sys"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func (s *Service) ChangePassword(ctx context.Context, accountID uuid.UUID, current, next, code string) error {
	emailAddr, _ := s.primaryContacts(ctx, accountID)
	if emailAddr == "" {
		return auth.ErrVerificationCodeWrong
	}
	if err := s.consumeVerificationCode(ctx, emailAddr, code); err != nil {
		return err
	}
	idents, err := s.repoFactory.Identity(none()).Get.ByAccountID(ctx, accountID)
	if err != nil {
		return sys.ErrNotFound
	}
	for _, item := range idents {
		if item.IdentityProvider() != "password" {
			continue
		}
		if item.PasswordHash() == nil || bcrypt.CompareHashAndPassword([]byte(*item.PasswordHash()), []byte(current)) != nil {
			return auth.ErrInvalidCredentials
		}
		hash, err := bcrypt.GenerateFromPassword([]byte(next), bcrypt.DefaultCost)
		if err != nil {
			return auth.ErrHashFailed
		}
		item.SetPasswordHash(string(hash))
		return s.repoFactory.Identity(none()).Update.Generic(ctx, item)
	}
	return sys.ErrNotFound
}

func (s *Service) SendPasswordCode(ctx context.Context, accountID uuid.UUID, lang string) error {
	emailAddr, _ := s.primaryContacts(ctx, accountID)
	if emailAddr == "" {
		return sys.ErrNotFound
	}
	return s.issueAndStoreVerification(ctx, emailAddr, lang)
}

func (s *Service) GetAccountByID(ctx context.Context, accountID uuid.UUID) (map[string]any, error) {
	return s.FullAccountWithProfile(ctx, accountID, uuid.Nil, "forger")
}

func (s *Service) FullAccountWithProfile(ctx context.Context, accountID uuid.UUID, profileID uuid.UUID, kind string) (map[string]any, error) {
	acc, err := s.repoFactory.Account(none()).Get.ByID(ctx, accountID)
	if err != nil {
		return nil, sys.ErrNotFound
	}
	emails, _ := s.emails.List.ByAccountID(ctx, accountID)
	phones, _ := s.phones.List.ByAccountID(ctx, accountID)
	idents, _ := s.repoFactory.Identity(none()).Get.ByAccountID(ctx, accountID)
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
		if p, err := s.repoFactory.Authority(none()).Get.ByAccountAndID(ctx, accountID, profileID); err == nil {
			m := s.authorityVO(p)
			s.attachProfileMedia(ctx, m, "authority", p.ID())
			out["authority_profile"] = m
		} else {
			out["authority_profile"] = nil
		}
	} else {
		if p, err := s.repoFactory.Forger(none()).Get.ByAccountAndID(ctx, accountID, profileID); err == nil {
			m := s.forgerVO(p)
			s.attachProfileMedia(ctx, m, "forger", p.ID())
			out["forger_profile"] = m
		} else {
			out["forger_profile"] = nil
		}
	}
	return out, nil
}

func (s *Service) attachProfileMedia(ctx context.Context, m map[string]any, kind string, profileID uuid.UUID) {
	if avatars, err := s.repoFactory.Avatar(none()).Get.ByProfileID(ctx, kind, profileID); err == nil {
		list := make([]map[string]any, 0, len(avatars))
		for _, a := range avatars {
			list = append(list, map[string]any{"id": a.ID().String(), "image_id": a.ImageID().String(), "is_active": a.IsActive()})
		}
		m["avatars"] = list
	}
	if bgs, err := s.repoFactory.Background(none()).Get.ByProfileID(ctx, kind, profileID); err == nil {
		list := make([]map[string]any, 0, len(bgs))
		for _, b := range bgs {
			list = append(list, map[string]any{"id": b.ID().String(), "image_id": b.ImageID().String(), "sort_order": b.SortOrder()})
		}
		m["backgrounds"] = list
	}
	if st, err := s.repoFactory.Settings(none()).Get.ByID(ctx, kind, profileID); err == nil {
		m["settings"] = map[string]any{"id": st.ID().String(), "login_notification": st.LoginNotification()}
	}
}
