package phone

import (
	"context"
	"time"

	"nfxidentity/errors/src/sys"
	phoneinternal "nfxidentity/modules/auth/application/phone/internal"

	"github.com/google/uuid"
)

type VerifyInput struct {
	AccountID uuid.UUID
	ProfileID uuid.UUID
	PhoneID   uuid.UUID
	Code      string
}

func (s *Service) Verify(ctx context.Context, in VerifyInput) error {
	row, err := s.repoFactory.Phone(phoneinternal.None()).Get.ByID(ctx, in.PhoneID)
	if err != nil || row.AccountID() != in.AccountID {
		return sys.ErrNotFound
	}
	if err := s.ConsumeVerificationCode(ctx, "phone:"+row.Phone(), in.Code); err != nil {
		return err
	}
	row.MarkVerified(time.Now())
	return s.repoFactory.Phone(phoneinternal.None()).Update.Generic(ctx, row)
}
