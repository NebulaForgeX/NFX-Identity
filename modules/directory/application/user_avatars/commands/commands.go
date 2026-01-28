package commands

import (
	"github.com/google/uuid"
)

// CreateOrUpdateUserAvatarCmd 创建或更新用户头像命令
type CreateOrUpdateUserAvatarCmd struct {
	UserID  uuid.UUID
	ImageID uuid.UUID
}

// UpdateUserAvatarImageIDCmd 更新用户头像图片ID命令
type UpdateUserAvatarImageIDCmd struct {
	UserID  uuid.UUID
	ImageID uuid.UUID
}

// DeleteUserAvatarCmd 删除用户头像命令
type DeleteUserAvatarCmd struct {
	UserID uuid.UUID
}
