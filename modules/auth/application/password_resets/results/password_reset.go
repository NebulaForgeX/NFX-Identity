package results

import (
	"time"

	"nfxid/modules/auth/domain/password_resets"

	"github.com/google/uuid"
)

type PasswordResetRO struct {
	ID          uuid.UUID
	ResetID     string
	TenantID    uuid.UUID
	UserID      uuid.UUID
	Delivery    password_resets.ResetDelivery
	CodeHash    string
	ExpiresAt   time.Time
	UsedAt      *time.Time
	RequestedIP *string
	UAHash      *string
	AttemptCount int
	Status      password_resets.ResetStatus
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// PasswordResetMapper 将 Domain PasswordReset 转换为 Application PasswordResetRO
func PasswordResetMapper(pr *password_resets.PasswordReset) PasswordResetRO {
	if pr == nil {
		return PasswordResetRO{}
	}

	return PasswordResetRO{
		ID:          pr.ID(),
		ResetID:     pr.ResetID(),
		TenantID:    pr.TenantID(),
		UserID:      pr.UserID(),
		Delivery:    pr.Delivery(),
		CodeHash:    pr.CodeHash(),
		ExpiresAt:   pr.ExpiresAt(),
		UsedAt:      pr.UsedAt(),
		RequestedIP:  pr.RequestedIP(),
		UAHash:      pr.UAHash(),
		AttemptCount: pr.AttemptCount(),
		Status:      pr.Status(),
		CreatedAt:   pr.CreatedAt(),
		UpdatedAt:   pr.UpdatedAt(),
	}
}
