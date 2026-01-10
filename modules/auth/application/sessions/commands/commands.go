package commands

import (
	"nfxid/modules/auth/domain/sessions"

	"github.com/google/uuid"
)

// CreateSessionCmd 创建会话命令
type CreateSessionCmd struct {
	SessionID         string
	TenantID          uuid.UUID
	UserID            uuid.UUID
	AppID             *uuid.UUID
	ClientID          *string
	ExpiresAt         string
	IP                *string
	UAHash            *string
	DeviceID          *string
	DeviceFingerprint *string
	DeviceName        *string
}

// UpdateSessionLastSeenCmd 更新会话最后访问时间命令
type UpdateSessionLastSeenCmd struct {
	SessionID string
}

// RevokeSessionCmd 撤销会话命令
type RevokeSessionCmd struct {
	SessionID    string
	RevokeReason sessions.SessionRevokeReason
	RevokedBy    string
}

// DeleteSessionCmd 删除会话命令
type DeleteSessionCmd struct {
	SessionID uuid.UUID
}
