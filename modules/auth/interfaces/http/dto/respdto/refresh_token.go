package respdto

import (
	"time"

	refreshTokenAppResult "nfxid/modules/auth/application/refresh_tokens/results"

	"github.com/google/uuid"
)

type RefreshTokenDTO struct {
	ID           uuid.UUID  `json:"id"`
	TokenID      string     `json:"token_id"`
	UserID       uuid.UUID  `json:"user_id"`
	TenantID     uuid.UUID  `json:"tenant_id"`
	AppID        *uuid.UUID `json:"app_id,omitempty"`
	ClientID     *string    `json:"client_id,omitempty"`
	SessionID    *uuid.UUID `json:"session_id,omitempty"`
	IssuedAt     time.Time  `json:"issued_at"`
	ExpiresAt    time.Time  `json:"expires_at"`
	RevokedAt    *time.Time `json:"revoked_at,omitempty"`
	RevokeReason *string    `json:"revoke_reason,omitempty"`
	RotatedFrom  *uuid.UUID `json:"rotated_from,omitempty"`
	DeviceID     *string    `json:"device_id,omitempty"`
	IP           *string    `json:"ip,omitempty"`
	UAHash       *string    `json:"ua_hash,omitempty"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
}

// RefreshTokenROToDTO converts application RefreshTokenRO to response DTO
func RefreshTokenROToDTO(v *refreshTokenAppResult.RefreshTokenRO) *RefreshTokenDTO {
	if v == nil {
		return nil
	}

	dto := &RefreshTokenDTO{
		ID:          v.ID,
		TokenID:     v.TokenID,
		UserID:      v.UserID,
		TenantID:    v.TenantID,
		AppID:       v.AppID,
		ClientID:    v.ClientID,
		SessionID:   v.SessionID,
		IssuedAt:    v.IssuedAt,
		ExpiresAt:   v.ExpiresAt,
		RevokedAt:   v.RevokedAt,
		RotatedFrom: v.RotatedFrom,
		DeviceID:    v.DeviceID,
		IP:          v.IP,
		UAHash:      v.UAHash,
		CreatedAt:   v.CreatedAt,
		UpdatedAt:   v.UpdatedAt,
	}

	if v.RevokeReason != nil {
		reason := string(*v.RevokeReason)
		dto.RevokeReason = &reason
	}

	return dto
}
