package commands

import (
	"github.com/google/uuid"
)

// CreateUserProfileCmd 创建用户资料命令
type CreateUserProfileCmd struct {
	UserID       uuid.UUID
	Role         *string
	FirstName    *string
	LastName     *string
	Nickname     *string
	DisplayName  *string
	AvatarID     *uuid.UUID
	BackgroundID *uuid.UUID
	BackgroundIDs []uuid.UUID
	Bio          *string
	Birthday     *string
	Age          *int
	Gender       *string
	Location     *string
	Website      *string
	Github       *string
	SocialLinks  map[string]interface{}
	Skills       map[string]interface{}
}

// UpdateUserProfileCmd 更新用户资料命令
type UpdateUserProfileCmd struct {
	UserProfileID uuid.UUID
	Role          *string
	FirstName     *string
	LastName      *string
	Nickname      *string
	DisplayName   *string
	AvatarID      *uuid.UUID
	BackgroundID *uuid.UUID
	BackgroundIDs []uuid.UUID
	Bio           *string
	Birthday      *string
	Age           *int
	Gender        *string
	Location      *string
	Website       *string
	Github        *string
	SocialLinks   map[string]interface{}
	Skills        map[string]interface{}
}

// DeleteUserProfileCmd 删除用户资料命令
type DeleteUserProfileCmd struct {
	UserProfileID uuid.UUID
}
