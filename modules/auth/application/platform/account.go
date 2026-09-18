package platform

import (
	"context"
	"nfxidentity/errors/src/auth"
	"nfxidentity/errors/src/sys"
	"nfxidentity/modules/auth/domain/profile"

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
		p, err := s.loadAuthorityProfile(ctx, accountID, profileID)
		if err == nil {
			m := s.authorityVO(p)
			s.attachAuthorityMedia(ctx, m, p.ID())
			out["authority_profile"] = m
		} else {
			out["authority_profile"] = nil
		}
	} else {
		p, err := s.loadForgerProfile(ctx, accountID, profileID)
		if err == nil {
			m := s.forgerVO(p)
			s.attachForgerMedia(ctx, m, p.ID())
			out["forger_profile"] = m
		} else {
			out["forger_profile"] = nil
		}
	}
	return out, nil
}

func (s *Service) loadForgerProfile(ctx context.Context, accountID, profileID uuid.UUID) (*profile.ForgerProfile, error) {
	if profileID == uuid.Nil {
		return s.repoFactory.Profile(none()).Get.ForgerByAccountID(ctx, accountID)
	}
	return s.repoFactory.Profile(none()).Get.ForgerByAccountAndID(ctx, accountID, profileID)
}

func (s *Service) loadAuthorityProfile(ctx context.Context, accountID, profileID uuid.UUID) (*profile.AuthorityProfile, error) {
	if profileID == uuid.Nil {
		return s.repoFactory.Profile(none()).Get.AuthorityByAccountID(ctx, accountID)
	}
	return s.repoFactory.Profile(none()).Get.AuthorityByAccountAndID(ctx, accountID, profileID)
}

func (s *Service) attachForgerMedia(ctx context.Context, m map[string]any, profileID uuid.UUID) {
	repo := s.repoFactory.Profile(none())
	if avatars, err := repo.Get.ListForgerAvatarsByProfileID(ctx, profileID); err == nil {
		list := make([]map[string]any, 0, len(avatars))
		for _, a := range avatars {
			list = append(list, map[string]any{"id": a.ID().String(), "image_id": a.ImageID().String(), "is_active": a.IsActive()})
		}
		m["avatars"] = list
	}
	if bgs, err := repo.Get.ForgerBackgroundsByProfileID(ctx, profileID); err == nil {
		list := make([]map[string]any, 0, len(bgs))
		for _, b := range bgs {
			list = append(list, map[string]any{"id": b.ID().String(), "image_id": b.ImageID().String(), "sort_order": b.SortOrder()})
		}
		m["backgrounds"] = list
	}
	if st, err := repo.Get.ForgerSettingsByProfileID(ctx, profileID); err == nil {
		m["settings"] = map[string]any{"id": st.ID().String(), "login_notification": st.LoginNotification()}
	}
}

func (s *Service) attachAuthorityMedia(ctx context.Context, m map[string]any, profileID uuid.UUID) {
	repo := s.repoFactory.Profile(none())
	if avatars, err := repo.Get.ListAuthorityAvatarsByProfileID(ctx, profileID); err == nil {
		list := make([]map[string]any, 0, len(avatars))
		for _, a := range avatars {
			list = append(list, map[string]any{"id": a.ID().String(), "image_id": a.ImageID().String(), "is_active": a.IsActive()})
		}
		m["avatars"] = list
	}
	if bgs, err := repo.Get.AuthorityBackgroundsByProfileID(ctx, profileID); err == nil {
		list := make([]map[string]any, 0, len(bgs))
		for _, b := range bgs {
			list = append(list, map[string]any{"id": b.ID().String(), "image_id": b.ImageID().String(), "sort_order": b.SortOrder()})
		}
		m["backgrounds"] = list
	}
	if st, err := repo.Get.AuthoritySettingsByProfileID(ctx, profileID); err == nil {
		m["settings"] = map[string]any{"id": st.ID().String(), "login_notification": st.LoginNotification()}
	}
}
