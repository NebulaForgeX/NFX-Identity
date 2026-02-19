package reqdto

import (
	sessionAppCommands "nfxid/modules/auth/application/sessions/commands"
	sessionDomain "nfxid/modules/auth/domain/sessions"

	"github.com/google/uuid"
)

type SessionCreateRequestDTO struct {
	SessionID         string     `json:"session_id" validate:"required"`
	TenantID          uuid.UUID  `json:"tenant_id" validate:"required"`
	UserID            uuid.UUID  `json:"user_id" validate:"required"`
	AppID             *uuid.UUID `json:"app_id,omitempty"`
	ClientID          *string    `json:"client_id,omitempty"`
	ExpiresAt         string     `json:"expires_at" validate:"required"`
	IP                *string    `json:"ip,omitempty"`
	UAHash            *string    `json:"ua_hash,omitempty"`
	DeviceID          *string    `json:"device_id,omitempty"`
	DeviceFingerprint *string    `json:"device_fingerprint,omitempty"`
	DeviceName        *string    `json:"device_name,omitempty"`
}

type SessionRevokeRequestDTO struct {
	SessionID    string `uri:"session_id" validate:"required"`
	RevokeReason string `json:"revoke_reason" validate:"required"`
	RevokedBy    string `json:"revoked_by" validate:"required"`
}

type SessionByIDRequestDTO struct {
	ID uuid.UUID `uri:"id" validate:"required,uuid"`
}

type SessionBySessionIDRequestDTO struct {
	SessionID string `uri:"session_id" validate:"required"`
}

func (r *SessionCreateRequestDTO) ToCreateCmd() sessionAppCommands.CreateSessionCmd {
	return sessionAppCommands.CreateSessionCmd{
		SessionID:         r.SessionID,
		TenantID:          r.TenantID,
		UserID:            r.UserID,
		AppID:             r.AppID,
		ClientID:          r.ClientID,
		ExpiresAt:         r.ExpiresAt,
		IP:                r.IP,
		UAHash:            r.UAHash,
		DeviceID:          r.DeviceID,
		DeviceFingerprint: r.DeviceFingerprint,
		DeviceName:        r.DeviceName,
	}
}

func (r *SessionRevokeRequestDTO) ToRevokeCmd() sessionAppCommands.RevokeSessionCmd {
	return sessionAppCommands.RevokeSessionCmd{
		SessionID:    r.SessionID,
		RevokeReason: sessionDomain.SessionRevokeReason(r.RevokeReason),
		RevokedBy:    r.RevokedBy,
	}
}
