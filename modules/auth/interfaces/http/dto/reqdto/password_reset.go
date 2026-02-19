package reqdto

import (
	passwordResetAppCommands "nfxid/modules/auth/application/password_resets/commands"
	passwordResetDomain "nfxid/modules/auth/domain/password_resets"

	"github.com/google/uuid"
)

type PasswordResetCreateRequestDTO struct {
	ResetID     string    `json:"reset_id" validate:"required"`
	TenantID    uuid.UUID `json:"tenant_id" validate:"required"`
	UserID      uuid.UUID `json:"user_id" validate:"required"`
	Delivery    string    `json:"delivery" validate:"required"`
	CodeHash    string    `json:"code_hash" validate:"required"`
	ExpiresAt   string    `json:"expires_at" validate:"required"`
	RequestedIP *string   `json:"requested_ip,omitempty"`
	UAHash      *string   `json:"ua_hash,omitempty"`
}

type PasswordResetUpdateRequestDTO struct {
	Status string `json:"status" validate:"required"`
}

type PasswordResetByIDRequestDTO struct {
	ID uuid.UUID `uri:"id" validate:"required,uuid"`
}

func (r *PasswordResetCreateRequestDTO) ToCreateCmd() passwordResetAppCommands.CreatePasswordResetCmd {
	return passwordResetAppCommands.CreatePasswordResetCmd{
		ResetID:     r.ResetID,
		TenantID:    r.TenantID,
		UserID:      r.UserID,
		Delivery:    passwordResetDomain.ResetDelivery(r.Delivery),
		CodeHash:    r.CodeHash,
		ExpiresAt:   r.ExpiresAt,
		RequestedIP: r.RequestedIP,
		UAHash:      r.UAHash,
	}
}

func (r *PasswordResetUpdateRequestDTO) ToUpdateStatusCmd(resetID string) passwordResetAppCommands.UpdateStatusCmd {
	return passwordResetAppCommands.UpdateStatusCmd{
		ResetID: resetID,
		Status:  passwordResetDomain.ResetStatus(r.Status),
	}
}
