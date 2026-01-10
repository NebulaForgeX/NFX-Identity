package commands

import (
	"nfxid/modules/auth/domain/password_resets"

	"github.com/google/uuid"
)

// CreatePasswordResetCmd 创建密码重置命令
type CreatePasswordResetCmd struct {
	ResetID     string
	TenantID    uuid.UUID
	UserID      uuid.UUID
	Delivery    password_resets.ResetDelivery
	CodeHash    string
	ExpiresAt   string
	RequestedIP *string
	UAHash      *string
}

// MarkAsUsedCmd 标记为已使用命令
type MarkAsUsedCmd struct {
	ResetID string
}

// IncrementAttemptCountCmd 增加尝试次数命令
type IncrementAttemptCountCmd struct {
	ResetID string
}

// UpdateStatusCmd 更新状态命令
type UpdateStatusCmd struct {
	ResetID string
	Status  password_resets.ResetStatus
}

// DeletePasswordResetCmd 删除密码重置命令
type DeletePasswordResetCmd struct {
	ResetID string
}
