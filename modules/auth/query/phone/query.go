package phone

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type PhoneItemVO struct {
	ID         uuid.UUID  `json:"id"`
	AccountID  uuid.UUID  `json:"account_id"`
	Phone      string     `json:"phone"`
	IsPrimary  bool       `json:"is_primary"`
	VerifiedAt *time.Time `json:"verified_at"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
}

type Query struct{ List List }

type List interface {
	ByAccountID(ctx context.Context, accountID uuid.UUID) ([]PhoneItemVO, error)
}
