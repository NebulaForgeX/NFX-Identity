package respdto

import (
	"time"

	sessionAppResult "nfxid/modules/auth/application/sessions/results"

	"github.com/google/uuid"
)

type SessionDTO struct {
	ID                uuid.UUID  `json:"id"`
	SessionID         string     `json:"session_id"`
	UserID            uuid.UUID  `json:"user_id"`
	AppID             *uuid.UUID `json:"app_id,omitempty"`
	ClientID          *string    `json:"client_id,omitempty"`
	CreatedAt         time.Time  `json:"created_at"`
	LastSeenAt        time.Time  `json:"last_seen_at"`
	ExpiresAt         time.Time  `json:"expires_at"`
	IP                *string    `json:"ip,omitempty"`
	UAHash            *string    `json:"ua_hash,omitempty"`
	DeviceID          *string    `json:"device_id,omitempty"`
	DeviceFingerprint *string    `json:"device_fingerprint,omitempty"`
	DeviceName        *string    `json:"device_name,omitempty"`
	RevokedAt         *time.Time `json:"revoked_at,omitempty"`
	RevokeReason      *string    `json:"revoke_reason,omitempty"`
	RevokedBy         *string    `json:"revoked_by,omitempty"`
	UpdatedAt         time.Time  `json:"updated_at"`
}

// SessionROToDTO converts application SessionRO to response DTO
func SessionROToDTO(v *sessionAppResult.SessionRO) *SessionDTO {
	if v == nil {
		return nil
	}

	dto := &SessionDTO{
		ID:                v.ID,
		SessionID:         v.SessionID,
		UserID:            v.UserID,
		AppID:             v.AppID,
		ClientID:          v.ClientID,
		CreatedAt:         v.CreatedAt,
		LastSeenAt:        v.LastSeenAt,
		ExpiresAt:         v.ExpiresAt,
		IP:                v.IP,
		UAHash:            v.UAHash,
		DeviceID:          v.DeviceID,
		DeviceFingerprint: v.DeviceFingerprint,
		DeviceName:        v.DeviceName,
		RevokedAt:         v.RevokedAt,
		RevokedBy:         v.RevokedBy,
		UpdatedAt:         v.UpdatedAt,
	}

	if v.RevokeReason != nil {
		reason := string(*v.RevokeReason)
		dto.RevokeReason = &reason
	}

	return dto
}
