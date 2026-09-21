package phone

import (
	"context"

	"nfxidentity/errors/src/sys"
	phoneinternal "nfxidentity/modules/auth/application/phone/internal"

	"github.com/google/uuid"
)

type SendVerificationCodeInput struct {
	AccountID uuid.UUID
	ProfileID uuid.UUID
	PhoneID   uuid.UUID
}

func (s *Service) SendVerificationCode(ctx context.Context, in SendVerificationCodeInput) error {
	row, err := s.repoFactory.Phone(phoneinternal.None()).Get.ByID(ctx, in.PhoneID)
	if err != nil || row.AccountID() != in.AccountID {
		return sys.ErrNotFound
	}
	return sys.ErrUnimplemented
}
