package password_resets

import "github.com/google/uuid"

func (pr *PasswordReset) Validate() error {
	if pr.ResetID() == "" {
		return ErrResetIDRequired
	}
	if pr.UserID() == uuid.Nil {
		return ErrUserIDRequired
	}
	if pr.TenantID() == uuid.Nil {
		return ErrTenantIDRequired
	}
	if pr.Delivery() == "" {
		return ErrDeliveryRequired
	}
	validDeliveries := map[ResetDelivery]struct{}{
		ResetDeliveryEmail: {},
		ResetDeliverySMS:   {},
	}
	if _, ok := validDeliveries[pr.Delivery()]; !ok {
		return ErrInvalidResetDelivery
	}
	if pr.CodeHash() == "" {
		return ErrCodeHashRequired
	}
	if pr.ExpiresAt().IsZero() {
		return ErrExpiresAtRequired
	}
	return nil
}
