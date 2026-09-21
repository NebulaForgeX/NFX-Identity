package account

import (
	"context"

	"nfxidentity/errors/src/sys"

	"github.com/google/uuid"
)

type SendChangePasswordVerificationCodeInput struct {
	AccountID uuid.UUID
	Lang      string
}

func (s *Service) SendChangePasswordVerificationCode(ctx context.Context, in SendChangePasswordVerificationCodeInput) error {
	emailAddr, _ := s.PrimaryContacts(ctx, in.AccountID)
	if emailAddr == "" {
		return sys.ErrNotFound
	}
	return s.IssueAndStoreVerification(ctx, emailAddr, in.Lang)
}
