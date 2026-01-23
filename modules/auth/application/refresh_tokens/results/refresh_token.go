package results

import (
	"time"

	"nfxid/modules/auth/domain/refresh_tokens"

	"github.com/google/uuid"
)

type RefreshTokenRO struct {
	ID           uuid.UUID
	TokenID      string
	UserID       uuid.UUID
	AppID        *uuid.UUID
	ClientID     *string
	SessionID    *uuid.UUID
	IssuedAt     time.Time
	ExpiresAt    time.Time
	RevokedAt    *time.Time
	RevokeReason *refresh_tokens.RevokeReason
	RotatedFrom  *uuid.UUID
	DeviceID     *string
	IP           *string
	UAHash       *string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

// RefreshTokenMapper 将 Domain RefreshToken 转换为 Application RefreshTokenRO
func RefreshTokenMapper(rt *refresh_tokens.RefreshToken) RefreshTokenRO {
	if rt == nil {
		return RefreshTokenRO{}
	}

	return RefreshTokenRO{
		ID:           rt.ID(),
		TokenID:      rt.TokenID(),
		UserID:       rt.UserID(),
		AppID:        rt.AppID(),
		ClientID:     rt.ClientID(),
		SessionID:    rt.SessionID(),
		IssuedAt:     rt.IssuedAt(),
		ExpiresAt:    rt.ExpiresAt(),
		RevokedAt:    rt.RevokedAt(),
		RevokeReason: rt.RevokeReason(),
		RotatedFrom:  rt.RotatedFrom(),
		DeviceID:     rt.DeviceID(),
		IP:           rt.IP(),
		UAHash:       rt.UAHash(),
		CreatedAt:    rt.CreatedAt(),
		UpdatedAt:    rt.UpdatedAt(),
	}
}
