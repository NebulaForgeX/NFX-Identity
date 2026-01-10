package commands

import (
	"github.com/google/uuid"
)

// CreateUserEmailCmd 创建用户邮箱命令
type CreateUserEmailCmd struct {
	UserID            uuid.UUID
	Email             string
	IsPrimary         bool
	IsVerified        bool
	VerificationToken *string
}

// SetPrimaryEmailCmd 设置主邮箱命令
type SetPrimaryEmailCmd struct {
	UserEmailID uuid.UUID
}

// VerifyEmailCmd 验证邮箱命令
type VerifyEmailCmd struct {
	UserEmailID uuid.UUID
}

// DeleteUserEmailCmd 删除用户邮箱命令
type DeleteUserEmailCmd struct {
	UserEmailID uuid.UUID
}
