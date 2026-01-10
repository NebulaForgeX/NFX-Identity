package commands

import (
	"github.com/google/uuid"
)

// CreatePasswordHistoryCmd 创建密码历史命令
type CreatePasswordHistoryCmd struct {
	UserID       uuid.UUID
	TenantID     uuid.UUID
	PasswordHash string
	HashAlg      *string
}

// DeletePasswordHistoryCmd 删除密码历史命令
type DeletePasswordHistoryCmd struct {
	PasswordHistoryID uuid.UUID
}
