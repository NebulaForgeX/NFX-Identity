package platform

import (
	"context"

	authmail "nfxidentity/modules/auth/infrastructure/email"
)

func (s *Service) issueAndStoreVerification(ctx context.Context, to, lang string) error {
	code := RandomCode()
	if err := authmail.SendVerificationEmail(ctx, s.mail, to, code, lang); err != nil {
		return err
	}
	return s.StoreVerificationCode(ctx, to, code)
}
