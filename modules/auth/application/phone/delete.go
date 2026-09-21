package phone

import (
	"context"
	"time"

	authErr "nfxidentity/errors/src/auth"
	phoneinternal "nfxidentity/modules/auth/application/phone/internal"

	"github.com/google/uuid"
)

type DeleteInput struct {
	AccountID uuid.UUID
	ProfileID uuid.UUID
	PhoneID   uuid.UUID
}

func (s *Service) Delete(ctx context.Context, in DeleteInput) error {
	row, err := s.repoFactory.Phone(phoneinternal.None()).Get.ByID(ctx, in.PhoneID)
	if err != nil || row.AccountID() != in.AccountID {
		return authErr.ErrPhoneNotDeletable
	}
	if row.IsPrimary() {
		return authErr.ErrPhoneNotDeletable
	}
	row.SoftDelete(time.Now())
	return s.repoFactory.Phone(phoneinternal.None()).Update.Generic(ctx, row)
}
