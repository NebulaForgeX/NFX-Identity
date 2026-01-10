package commands

import (
	"github.com/google/uuid"
)

// CreateAPIKeyCmd 创建API密钥命令
type CreateAPIKeyCmd struct {
	KeyID       string
	AppID       uuid.UUID
	KeyHash     string
	HashAlg     string
	Name        string
	ExpiresAt   *string
	CreatedBy   *uuid.UUID
	Metadata    map[string]interface{}
}

// RevokeAPIKeyCmd 撤销API密钥命令
type RevokeAPIKeyCmd struct {
	KeyID        string
	RevokedBy    uuid.UUID
	RevokeReason *string
}

// UpdateLastUsedCmd 更新最后使用时间命令
type UpdateLastUsedCmd struct {
	KeyID string
}

// DeleteAPIKeyCmd 删除API密钥命令
type DeleteAPIKeyCmd struct {
	KeyID string
}
