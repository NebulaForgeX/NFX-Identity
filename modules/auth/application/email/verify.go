package email

import (
	"context"
	"time"

	"nfxidentity/errors/src/sys"
	emailinternal "nfxidentity/modules/auth/application/email/internal"

	"github.com/google/uuid"
)

type VerifyInput struct {
	AccountID uuid.UUID
	ProfileID uuid.UUID
	EmailID   uuid.UUID
	Code      string
}

func (s *Service) Verify(ctx context.Context, in VerifyInput) error {
	row, err := s.repoFactory.Email(emailinternal.None()).Get.ByID(ctx, in.EmailID)
	if err != nil || row.AccountID() != in.AccountID {
		return sys.ErrNotFound
	}
	if err := s.ConsumeVerificationCode(ctx, row.Email(), in.Code); err != nil {
		return err
	}
	row.MarkVerified(time.Now())
	return s.repoFactory.Email(emailinternal.None()).Update.Generic(ctx, row)
}
