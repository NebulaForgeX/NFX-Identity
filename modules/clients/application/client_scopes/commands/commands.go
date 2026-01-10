package commands

import (
	"github.com/google/uuid"
)

// CreateClientScopeCmd 创建客户端作用域命令
type CreateClientScopeCmd struct {
	AppID      uuid.UUID
	Scope      string
	GrantedBy  *uuid.UUID
	ExpiresAt  *string
}

// RevokeClientScopeCmd 撤销客户端作用域命令
type RevokeClientScopeCmd struct {
	ClientScopeID uuid.UUID
	RevokedBy     uuid.UUID
	RevokeReason  *string
}

// DeleteClientScopeCmd 删除客户端作用域命令
type DeleteClientScopeCmd struct {
	ClientScopeID uuid.UUID
}
