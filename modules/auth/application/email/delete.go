package email

import (
	"context"
	"time"

	authErr "nfxidentity/errors/src/auth"
	emailinternal "nfxidentity/modules/auth/application/email/internal"

	"github.com/google/uuid"
)

type DeleteInput struct {
	AccountID uuid.UUID
	ProfileID uuid.UUID
	EmailID   uuid.UUID
}

func (s *Service) Delete(ctx context.Context, in DeleteInput) error {
	row, err := s.repoFactory.Email(emailinternal.None()).Get.ByID(ctx, in.EmailID)
	if err != nil || row.AccountID() != in.AccountID {
		return authErr.ErrEmailCannotDeletePrimary
	}
	if row.IsPrimary() {
		return authErr.ErrEmailCannotDeletePrimary
	}
	row.SoftDelete(time.Now())
	return s.repoFactory.Email(emailinternal.None()).Update.Generic(ctx, row)
}
