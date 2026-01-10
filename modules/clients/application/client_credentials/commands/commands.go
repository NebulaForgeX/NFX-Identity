package commands

import (
	"github.com/google/uuid"
)

// CreateClientCredentialCmd 创建客户端凭证命令
type CreateClientCredentialCmd struct {
	AppID      uuid.UUID
	ClientID   string
	SecretHash string
	HashAlg    string
	ExpiresAt  *string
	CreatedBy  *uuid.UUID
}

// RevokeClientCredentialCmd 撤销客户端凭证命令
type RevokeClientCredentialCmd struct {
	ClientID     string
	RevokedBy    uuid.UUID
	RevokeReason *string
}

// RotateClientCredentialCmd 轮换客户端凭证命令
type RotateClientCredentialCmd struct {
	ClientID     string
	NewSecretHash string
	NewHashAlg   string
}

// UpdateLastUsedCmd 更新最后使用时间命令
type UpdateLastUsedCmd struct {
	ClientID string
}

// DeleteClientCredentialCmd 删除客户端凭证命令
type DeleteClientCredentialCmd struct {
	ClientID string
}
