package email

import (
	"context"
	"strings"

	"nfxidentity/errors/src/sys"
	emailinternal "nfxidentity/modules/auth/application/email/internal"

	"github.com/google/uuid"
)

type UpdateInput struct {
	AccountID uuid.UUID
	ProfileID uuid.UUID
	EmailID   uuid.UUID
	Email     string
}

func (s *Service) Update(ctx context.Context, in UpdateInput) error {
	address := strings.ToLower(strings.TrimSpace(in.Email))
	row, err := s.repoFactory.Email(emailinternal.None()).Get.ByID(ctx, in.EmailID)
	if err != nil || row.AccountID() != in.AccountID {
		return sys.ErrNotFound
	}
	row.ChangeAddress(address)
	return s.repoFactory.Email(emailinternal.None()).Update.Generic(ctx, row)
}
