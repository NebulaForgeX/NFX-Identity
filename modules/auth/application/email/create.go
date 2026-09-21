package email

import (
	"context"
	"strings"
	"time"

	authErr "nfxidentity/errors/src/auth"
	emailinternal "nfxidentity/modules/auth/application/email/internal"
	emaildomain "nfxidentity/modules/auth/domain/email"

	"github.com/google/uuid"
)

type CreateInput struct {
	AccountID uuid.UUID
	ProfileID uuid.UUID
	Email     string
}

type CreateOutput struct {
	EmailID uuid.UUID `json:"email_id"`
}

func (s *Service) Create(ctx context.Context, in CreateInput) (*CreateOutput, error) {
	address := strings.ToLower(strings.TrimSpace(in.Email))
	if address == "" {
		return nil, authErr.ErrInvalidEmail
	}
	now := time.Now()
	id := uuid.New()
	if err := s.repoFactory.Email(emailinternal.None()).Create.New(ctx, emaildomain.NewEmailFromState(emaildomain.EmailState{
		ID: id, AccountID: in.AccountID, Email: address, CreatedAt: now, UpdatedAt: now,
	})); err != nil {
		return nil, authErr.ErrEmailAlreadyExists
	}
	return &CreateOutput{EmailID: id}, nil
}
