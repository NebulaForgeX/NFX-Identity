package email

import (
	"context"

	"nfxidentity/errors/src/sys"
	emailinternal "nfxidentity/modules/auth/application/email/internal"

	"github.com/google/uuid"
)

type SendVerificationCodeInput struct {
	AccountID uuid.UUID
	ProfileID uuid.UUID
	EmailID   uuid.UUID
	Lang      string
}

func (s *Service) SendVerificationCode(ctx context.Context, in SendVerificationCodeInput) error {
	row, err := s.repoFactory.Email(emailinternal.None()).Get.ByID(ctx, in.EmailID)
	if err != nil || row.AccountID() != in.AccountID {
		return sys.ErrNotFound
	}
	return s.IssueAndStoreVerification(ctx, row.Email(), in.Lang)
}
