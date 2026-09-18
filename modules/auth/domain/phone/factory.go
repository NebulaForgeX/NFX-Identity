package phone

import (
	"strings"
	"time"

	authErr "nfxidentity/errors/src/auth"

	"github.com/google/uuid"
)

type NewPhoneParams struct {
	AccountID uuid.UUID
	Phone     string
	IsPrimary bool
}

func NewPhone(p NewPhoneParams) (*Phone, error) {
	if p.AccountID == uuid.Nil {
		return nil, authErr.ErrPhoneAccountIDInvalid
	}
	num := strings.TrimSpace(p.Phone)
	if err := validatePhoneNumber(num); err != nil {
		return nil, err
	}
	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}
	now := time.Now().UTC()
	return NewPhoneFromState(PhoneState{
		ID:        id,
		AccountID: p.AccountID,
		Phone:     num,
		IsPrimary: p.IsPrimary,
		CreatedAt: now,
		UpdatedAt: now,
	}), nil
}
