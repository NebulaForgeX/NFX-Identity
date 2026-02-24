package password_resets

import (
	authErr "nfxid/errors/src/auth"

	"github.com/google/uuid"
)

func (pr *PasswordReset) Validate() error {
	if pr.ResetID() == "" {
		return authErr.ErrResetIDRequired
	}
	if pr.UserID() == uuid.Nil {
		return authErr.ErrUserIDRequired
	}
	if pr.TenantID() == uuid.Nil {
		return authErr.ErrTenantIDRequired
	}
	if pr.Delivery() == "" {
		return authErr.ErrDeliveryRequired
	}
	validDeliveries := map[ResetDelivery]struct{}{
		ResetDeliveryEmail: {},
		ResetDeliverySMS:   {},
	}
	if _, ok := validDeliveries[pr.Delivery()]; !ok {
		return authErr.ErrInvalidResetDelivery
	}
	if pr.CodeHash() == "" {
		return authErr.ErrCodeHashRequired
	}
	if pr.ExpiresAt().IsZero() {
		return authErr.ErrExpiresAtRequired
	}
	return nil
}
