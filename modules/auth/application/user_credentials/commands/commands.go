package commands

import (
	"nfxid/modules/auth/domain/user_credentials"

	"github.com/google/uuid"
)

// CreateUserCredentialCmd 创建用户凭证命令
type CreateUserCredentialCmd struct {
	UserID            uuid.UUID
	CredentialType    user_credentials.CredentialType
	PasswordHash      *string
	HashAlg           *string
	HashParams        map[string]interface{}
	Status            user_credentials.CredentialStatus
	MustChangePassword bool
}

// UpdateUserCredentialCmd 更新用户凭证命令
type UpdateUserCredentialCmd struct {
	UserCredentialID uuid.UUID
	Status           user_credentials.CredentialStatus
	MustChangePassword bool
}

// UpdatePasswordCmd 更新密码命令
type UpdatePasswordCmd struct {
	UserID       uuid.UUID
	PasswordHash string
	HashAlg      string
	HashParams   map[string]interface{}
}

// UpdateStatusCmd 更新状态命令
type UpdateStatusCmd struct {
	UserID uuid.UUID
	Status user_credentials.CredentialStatus
}

// DeleteUserCredentialCmd 删除用户凭证命令
type DeleteUserCredentialCmd struct {
	UserCredentialID uuid.UUID
}
