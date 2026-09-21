package phone

import (
	"context"
	"strings"

	"nfxidentity/errors/src/sys"
	phoneinternal "nfxidentity/modules/auth/application/phone/internal"

	"github.com/google/uuid"
)

type UpdateInput struct {
	AccountID uuid.UUID
	ProfileID uuid.UUID
	PhoneID   uuid.UUID
	Phone     string
}

func (s *Service) Update(ctx context.Context, in UpdateInput) error {
	number := strings.TrimSpace(in.Phone)
	row, err := s.repoFactory.Phone(phoneinternal.None()).Get.ByID(ctx, in.PhoneID)
	if err != nil || row.AccountID() != in.AccountID {
		return sys.ErrNotFound
	}
	row.ChangeNumber(number)
	return s.repoFactory.Phone(phoneinternal.None()).Update.Generic(ctx, row)
}
