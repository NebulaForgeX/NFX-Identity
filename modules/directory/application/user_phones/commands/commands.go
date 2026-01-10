package commands

import (
	"github.com/google/uuid"
)

// CreateUserPhoneCmd 创建用户手机号命令
type CreateUserPhoneCmd struct {
	UserID                uuid.UUID
	Phone                 string
	CountryCode           *string
	IsPrimary             bool
	IsVerified            bool
	VerificationCode      *string
	VerificationExpiresAt *string
}

// SetPrimaryPhoneCmd 设置主手机号命令
type SetPrimaryPhoneCmd struct {
	UserPhoneID uuid.UUID
}

// VerifyPhoneCmd 验证手机号命令
type VerifyPhoneCmd struct {
	UserPhoneID uuid.UUID
}

// UpdateVerificationCodeCmd 更新验证码命令
type UpdateVerificationCodeCmd struct {
	UserPhoneID           uuid.UUID
	VerificationCode      string
	VerificationExpiresAt string
}

// DeleteUserPhoneCmd 删除用户手机号命令
type DeleteUserPhoneCmd struct {
	UserPhoneID uuid.UUID
}
