package reqdto

import (
	refreshTokenAppCommands "nfxid/modules/auth/application/refresh_tokens/commands"
	refreshTokenDomain "nfxid/modules/auth/domain/refresh_tokens"

	"github.com/google/uuid"
)

type RefreshTokenCreateRequestDTO struct {
	TokenID   string     `json:"token_id" validate:"required"`
	UserID    uuid.UUID  `json:"user_id" validate:"required"`
	AppID     *uuid.UUID `json:"app_id,omitempty"`
	ClientID  *string    `json:"client_id,omitempty"`
	SessionID *uuid.UUID `json:"session_id,omitempty"`
	ExpiresAt string     `json:"expires_at" validate:"required"`
	DeviceID  *string    `json:"device_id,omitempty"`
	IP        *string    `json:"ip,omitempty"`
	UAHash    *string    `json:"ua_hash,omitempty"`
}

type RefreshTokenRevokeRequestDTO struct {
	RevokeReason string `json:"revoke_reason" validate:"required"`
}

type RefreshTokenByIDRequestDTO struct {
	ID uuid.UUID `params:"id" validate:"required,uuid"`
}

func (r *RefreshTokenCreateRequestDTO) ToCreateCmd() refreshTokenAppCommands.CreateRefreshTokenCmd {
	return refreshTokenAppCommands.CreateRefreshTokenCmd{
		TokenID:   r.TokenID,
		UserID:    r.UserID,
		AppID:     r.AppID,
		ClientID:  r.ClientID,
		SessionID: r.SessionID,
		ExpiresAt: r.ExpiresAt,
		DeviceID:  r.DeviceID,
		IP:        r.IP,
		UAHash:    r.UAHash,
	}
}

func (r *RefreshTokenRevokeRequestDTO) ToRevokeCmd(tokenID string) refreshTokenAppCommands.RevokeRefreshTokenCmd {
	return refreshTokenAppCommands.RevokeRefreshTokenCmd{
		TokenID:      tokenID,
		RevokeReason: refreshTokenDomain.RevokeReason(r.RevokeReason),
	}
}
