package commands

import (
	"nfxid/modules/auth/domain/refresh_tokens"

	"github.com/google/uuid"
)

// CreateRefreshTokenCmd 创建刷新令牌命令
type CreateRefreshTokenCmd struct {
	TokenID   string
	UserID    uuid.UUID
	TenantID  uuid.UUID
	AppID     *uuid.UUID
	ClientID  *string
	SessionID *uuid.UUID
	ExpiresAt string
	DeviceID  *string
	IP        *string
	UAHash    *string
}

// RevokeRefreshTokenCmd 撤销刷新令牌命令
type RevokeRefreshTokenCmd struct {
	TokenID      string
	RevokeReason refresh_tokens.RevokeReason
}

// DeleteRefreshTokenCmd 删除刷新令牌命令
type DeleteRefreshTokenCmd struct {
	RefreshTokenID uuid.UUID
}
