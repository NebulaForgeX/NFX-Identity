package commands

import (
	"nfxid/modules/directory/domain/users"

	"github.com/google/uuid"
)

// CreateUserCmd 创建用户命令
type CreateUserCmd struct {
	TenantID   uuid.UUID
	Username  string
	Status    users.UserStatus
	IsVerified bool
}

// UpdateUserStatusCmd 更新用户状态命令
type UpdateUserStatusCmd struct {
	UserID uuid.UUID
	Status users.UserStatus
}

// UpdateUsernameCmd 更新用户名命令
type UpdateUsernameCmd struct {
	UserID   uuid.UUID
	Username string
}

// VerifyUserCmd 验证用户命令
type VerifyUserCmd struct {
	UserID uuid.UUID
}

// DeleteUserCmd 删除用户命令
type DeleteUserCmd struct {
	UserID uuid.UUID
}
