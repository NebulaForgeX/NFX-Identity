package sessions

import (
	"time"

	"github.com/google/uuid"
)

type NewSessionParams struct {
	SessionID        string
	TenantID         uuid.UUID
	UserID           uuid.UUID
	AppID            *uuid.UUID
	ClientID         *string
	ExpiresAt        time.Time
	IP               *string
	UAHash           *string
	DeviceID         *string
	DeviceFingerprint *string
	DeviceName       *string
}

func NewSession(p NewSessionParams) (*Session, error) {
	if err := validateSessionParams(p); err != nil {
		return nil, err
	}

	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	now := time.Now().UTC()
	return NewSessionFromState(SessionState{
		ID:                id,
		SessionID:         p.SessionID,
		TenantID:          p.TenantID,
		UserID:            p.UserID,
		AppID:             p.AppID,
		ClientID:          p.ClientID,
		CreatedAt:         now,
		LastSeenAt:        now,
		ExpiresAt:         p.ExpiresAt,
		IP:                p.IP,
		UAHash:            p.UAHash,
		DeviceID:          p.DeviceID,
		DeviceFingerprint: p.DeviceFingerprint,
		DeviceName:        p.DeviceName,
		UpdatedAt:         now,
	}), nil
}

func NewSessionFromState(st SessionState) *Session {
	return &Session{state: st}
}

func validateSessionParams(p NewSessionParams) error {
	if p.SessionID == "" {
		return ErrSessionIDRequired
	}
	if p.UserID == uuid.Nil {
		return ErrUserIDRequired
	}
	if p.TenantID == uuid.Nil {
		return ErrTenantIDRequired
	}
	if p.ExpiresAt.IsZero() {
		return ErrExpiresAtRequired
	}
	return nil
}
