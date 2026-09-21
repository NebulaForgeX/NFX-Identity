package signup

import (
	"context"

	authErr "nfxidentity/errors/src/auth"
	signupinternal "nfxidentity/modules/auth/application/signup/internal"
	authmail "nfxidentity/modules/auth/infrastructure/email"
)

const verificationEntityPurpose = "signup"

func (s *Service) BySendingCode(ctx context.Context, email, lang string) error {
	if registered, err := s.repoFactory.Email(signupinternal.None()).Check.ByEmail(ctx, email); err != nil {
		return authErr.ErrEmailRegistrationCheckFailed.WithCause(err)
	} else if registered {
		return authErr.ErrEmailAlreadyExists
	}
	code := signupinternal.RandomCode()
	if err := authmail.SendVerificationEmail(ctx, s.mail, email, code, lang); err != nil {
		return authErr.ErrSendVerificationEmailFailed.WithCause(err)
	}
	if err := s.storeVerificationCode(ctx, email, code); err != nil {
		return authErr.ErrVerificationCodeSaveFailed.WithCause(err)
	}
	return nil
}
