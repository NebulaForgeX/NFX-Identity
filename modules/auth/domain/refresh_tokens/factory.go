package refresh_tokens

import (
	"time"

	"github.com/google/uuid"
)

type NewRefreshTokenParams struct {
	TokenID     string
	UserID      uuid.UUID
	AppID       *uuid.UUID
	ClientID    *string
	SessionID   *uuid.UUID
	ExpiresAt   time.Time
	RotatedFrom *uuid.UUID
	DeviceID    *string
	IP          *string
	UAHash      *string
}

func NewRefreshToken(p NewRefreshTokenParams) (*RefreshToken, error) {
	if err := validateRefreshTokenParams(p); err != nil {
		return nil, err
	}

	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	now := time.Now().UTC()
	return NewRefreshTokenFromState(RefreshTokenState{
		ID:          id,
		TokenID:     p.TokenID,
		UserID:      p.UserID,
		AppID:       p.AppID,
		ClientID:    p.ClientID,
		SessionID:   p.SessionID,
		IssuedAt:    now,
		ExpiresAt:   p.ExpiresAt,
		RotatedFrom: p.RotatedFrom,
		DeviceID:    p.DeviceID,
		IP:          p.IP,
		UAHash:      p.UAHash,
		CreatedAt:   now,
		UpdatedAt:   now,
	}), nil
}

func NewRefreshTokenFromState(st RefreshTokenState) *RefreshToken {
	return &RefreshToken{state: st}
}

func validateRefreshTokenParams(p NewRefreshTokenParams) error {
	if p.TokenID == "" {
		return ErrTokenIDRequired
	}
	if p.UserID == uuid.Nil {
		return ErrUserIDRequired
	}
	if p.ExpiresAt.IsZero() {
		return ErrExpiresAtRequired
	}
	return nil
}
