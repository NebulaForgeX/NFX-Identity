package account

import (
	"context"

	"nfxidentity/errors/src/auth"
	"nfxidentity/errors/src/sys"
	accountinternal "nfxidentity/modules/auth/application/account/internal"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type ChangePasswordInput struct {
	AccountID        uuid.UUID
	CurrentPassword  string
	NewPassword      string
	VerificationCode string
}

func (s *Service) ChangePassword(ctx context.Context, in ChangePasswordInput) error {
	emailAddr, _ := s.PrimaryContacts(ctx, in.AccountID)
	if emailAddr == "" {
		return auth.ErrVerificationCodeWrong
	}
	if err := s.ConsumeVerificationCode(ctx, emailAddr, in.VerificationCode); err != nil {
		return err
	}
	idents, err := s.repoFactory.Identity(accountinternal.None()).Get.ByAccountID(ctx, in.AccountID)
	if err != nil {
		return sys.ErrNotFound
	}
	for _, item := range idents {
		if item.IdentityProvider() != "password" {
			continue
		}
		if item.PasswordHash() == nil || bcrypt.CompareHashAndPassword([]byte(*item.PasswordHash()), []byte(in.CurrentPassword)) != nil {
			return auth.ErrInvalidCredentials
		}
		hash, err := bcrypt.GenerateFromPassword([]byte(in.NewPassword), bcrypt.DefaultCost)
		if err != nil {
			return auth.ErrHashFailed
		}
		item.SetPasswordHash(string(hash))
		return s.repoFactory.Identity(accountinternal.None()).Update.Generic(ctx, item)
	}
	return sys.ErrNotFound
}
