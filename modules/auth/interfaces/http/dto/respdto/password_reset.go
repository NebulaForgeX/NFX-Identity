package respdto

import (
	"time"

	passwordResetAppResult "nfxid/modules/auth/application/password_resets/results"

	"github.com/google/uuid"
)

type PasswordResetDTO struct {
	ID          uuid.UUID  `json:"id"`
	ResetID     string     `json:"reset_id"`
	TenantID    uuid.UUID  `json:"tenant_id"`
	UserID      uuid.UUID  `json:"user_id"`
	Delivery    string     `json:"delivery"`
	CodeHash    string     `json:"code_hash"`
	ExpiresAt   time.Time  `json:"expires_at"`
	UsedAt      *time.Time `json:"used_at,omitempty"`
	RequestedIP *string    `json:"requested_ip,omitempty"`
	UAHash      *string    `json:"ua_hash,omitempty"`
	AttemptCount int       `json:"attempt_count"`
	Status      string     `json:"status"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

// PasswordResetROToDTO converts application PasswordResetRO to response DTO
func PasswordResetROToDTO(v *passwordResetAppResult.PasswordResetRO) *PasswordResetDTO {
	if v == nil {
		return nil
	}

	return &PasswordResetDTO{
		ID:          v.ID,
		ResetID:     v.ResetID,
		TenantID:    v.TenantID,
		UserID:      v.UserID,
		Delivery:    string(v.Delivery),
		CodeHash:    v.CodeHash,
		ExpiresAt:   v.ExpiresAt,
		UsedAt:      v.UsedAt,
		RequestedIP: v.RequestedIP,
		UAHash:      v.UAHash,
		AttemptCount: v.AttemptCount,
		Status:      string(v.Status),
		CreatedAt:   v.CreatedAt,
		UpdatedAt:   v.UpdatedAt,
	}
}
