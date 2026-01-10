package results

import (
	"time"

	"nfxid/modules/auth/domain/sessions"

	"github.com/google/uuid"
)

type SessionRO struct {
	ID                uuid.UUID
	SessionID         string
	TenantID          uuid.UUID
	UserID            uuid.UUID
	AppID             *uuid.UUID
	ClientID          *string
	CreatedAt         time.Time
	LastSeenAt        time.Time
	ExpiresAt         time.Time
	IP                *string
	UAHash            *string
	DeviceID          *string
	DeviceFingerprint *string
	DeviceName        *string
	RevokedAt         *time.Time
	RevokeReason      *sessions.SessionRevokeReason
	RevokedBy         *string
	UpdatedAt         time.Time
}

// SessionMapper 将 Domain Session 转换为 Application SessionRO
func SessionMapper(s *sessions.Session) SessionRO {
	if s == nil {
		return SessionRO{}
	}

	return SessionRO{
		ID:                s.ID(),
		SessionID:         s.SessionID(),
		TenantID:          s.TenantID(),
		UserID:            s.UserID(),
		AppID:             s.AppID(),
		ClientID:          s.ClientID(),
		CreatedAt:         s.CreatedAt(),
		LastSeenAt:        s.LastSeenAt(),
		ExpiresAt:         s.ExpiresAt(),
		IP:                s.IP(),
		UAHash:            s.UAHash(),
		DeviceID:          s.DeviceID(),
		DeviceFingerprint: s.DeviceFingerprint(),
		DeviceName:        s.DeviceName(),
		RevokedAt:         s.RevokedAt(),
		RevokeReason:      s.RevokeReason(),
		RevokedBy:         s.RevokedBy(),
		UpdatedAt:         s.UpdatedAt(),
	}
}
