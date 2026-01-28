package commands

import (
	"github.com/google/uuid"
)

// CreateUserImageCmd 创建用户图片命令
type CreateUserImageCmd struct {
	UserID       uuid.UUID
	ImageID      uuid.UUID
	DisplayOrder int
}

// UpdateUserImageDisplayOrderCmd 更新用户图片显示顺序命令
type UpdateUserImageDisplayOrderCmd struct {
	UserImageID uuid.UUID
	DisplayOrder int
}

// UpdateUserImageImageIDCmd 更新用户图片ID命令
type UpdateUserImageImageIDCmd struct {
	UserImageID uuid.UUID
	ImageID     uuid.UUID
}

// DeleteUserImageCmd 删除用户图片命令
type DeleteUserImageCmd struct {
	UserImageID uuid.UUID
}
