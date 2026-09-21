package phone

import (
	"context"
	"strings"
	"time"

	authErr "nfxidentity/errors/src/auth"
	phoneinternal "nfxidentity/modules/auth/application/phone/internal"
	phonedomain "nfxidentity/modules/auth/domain/phone"

	"github.com/google/uuid"
)

type CreateInput struct {
	AccountID uuid.UUID
	ProfileID uuid.UUID
	Phone     string
}

type CreateOutput struct {
	PhoneID uuid.UUID `json:"phone_id"`
}

func (s *Service) Create(ctx context.Context, in CreateInput) (*CreateOutput, error) {
	number := strings.TrimSpace(in.Phone)
	if number == "" {
		return nil, authErr.ErrInvalidPhone
	}
	now := time.Now()
	id := uuid.New()
	if err := s.repoFactory.Phone(phoneinternal.None()).Create.New(ctx, phonedomain.NewPhoneFromState(phonedomain.PhoneState{
		ID: id, AccountID: in.AccountID, Phone: number, CreatedAt: now, UpdatedAt: now,
	})); err != nil {
		return nil, authErr.ErrPhoneAlreadyExists
	}
	return &CreateOutput{PhoneID: id}, nil
}
