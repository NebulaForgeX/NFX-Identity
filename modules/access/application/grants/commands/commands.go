package commands

import (
	"nfxid/modules/access/domain/grants"

	"github.com/google/uuid"
)

// CreateGrantCmd 创建授权命令
type CreateGrantCmd struct {
	SubjectType  grants.SubjectType
	SubjectID    uuid.UUID
	GrantType    grants.GrantType
	GrantRefID   uuid.UUID
	TenantID     *uuid.UUID
	AppID        *uuid.UUID
	ResourceType *string
	ResourceID   *uuid.UUID
	Effect       grants.GrantEffect
	ExpiresAt    *string
	CreatedBy    *uuid.UUID
}

// UpdateGrantCmd 更新授权命令
type UpdateGrantCmd struct {
	GrantID      uuid.UUID
	ExpiresAt    *string
}

// RevokeGrantCmd 撤销授权命令
type RevokeGrantCmd struct {
	GrantID      uuid.UUID
	RevokedBy    uuid.UUID
	RevokeReason *string
}

// DeleteGrantCmd 删除授权命令
type DeleteGrantCmd struct {
	GrantID uuid.UUID
}
