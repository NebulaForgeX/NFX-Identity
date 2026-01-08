package password_resets

import (
	"time"

	"github.com/google/uuid"
)

type NewPasswordResetParams struct {
	ResetID     string
	TenantID    uuid.UUID
	UserID      uuid.UUID
	Delivery    ResetDelivery
	CodeHash    string
	ExpiresAt   time.Time
	RequestedIP *string
	UAHash      *string
}

func NewPasswordReset(p NewPasswordResetParams) (*PasswordReset, error) {
	if err := validatePasswordResetParams(p); err != nil {
		return nil, err
	}

	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	now := time.Now().UTC()
	return NewPasswordResetFromState(PasswordResetState{
		ID:           id,
		ResetID:      p.ResetID,
		TenantID:     p.TenantID,
		UserID:       p.UserID,
		Delivery:     p.Delivery,
		CodeHash:     p.CodeHash,
		ExpiresAt:    p.ExpiresAt,
		RequestedIP:  p.RequestedIP,
		UAHash:       p.UAHash,
		AttemptCount: 0,
		Status:       ResetStatusIssued,
		CreatedAt:    now,
		UpdatedAt:    now,
	}), nil
}

func NewPasswordResetFromState(st PasswordResetState) *PasswordReset {
	return &PasswordReset{state: st}
}

func validatePasswordResetParams(p NewPasswordResetParams) error {
	if p.ResetID == "" {
		return ErrResetIDRequired
	}
	if p.UserID == uuid.Nil {
		return ErrUserIDRequired
	}
	if p.TenantID == uuid.Nil {
		return ErrTenantIDRequired
	}
	if p.Delivery == "" {
		return ErrDeliveryRequired
	}
	validDeliveries := map[ResetDelivery]struct{}{
		ResetDeliveryEmail: {},
		ResetDeliverySMS:   {},
	}
	if _, ok := validDeliveries[p.Delivery]; !ok {
		return ErrInvalidResetDelivery
	}
	if p.CodeHash == "" {
		return ErrCodeHashRequired
	}
	if p.ExpiresAt.IsZero() {
		return ErrExpiresAtRequired
	}
	return nil
}
