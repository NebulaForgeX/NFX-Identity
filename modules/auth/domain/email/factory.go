package email

import (
	"strings"
	"time"

	authErr "nfxidentity/errors/src/auth"

	"github.com/google/uuid"
)

type NewEmailParams struct {
	AccountID uuid.UUID
	Email     string
	IsPrimary bool
}

func NewEmail(p NewEmailParams) (*Email, error) {
	if p.AccountID == uuid.Nil {
		return nil, authErr.ErrEmailAccountIDInvalid
	}
	addr := strings.TrimSpace(p.Email)
	if err := validateEmailAddress(addr); err != nil {
		return nil, err
	}
	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}
	now := time.Now().UTC()
	return NewEmailFromState(EmailState{
		ID:        id,
		AccountID: p.AccountID,
		Email:     addr,
		IsPrimary: p.IsPrimary,
		CreatedAt: now,
		UpdatedAt: now,
	}), nil
}
