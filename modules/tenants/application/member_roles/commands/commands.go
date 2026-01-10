package commands

import (
	"github.com/google/uuid"
)

// CreateMemberRoleCmd 创建成员角色命令
type CreateMemberRoleCmd struct {
	TenantID   uuid.UUID
	MemberID   uuid.UUID
	RoleID     uuid.UUID
	AssignedBy *uuid.UUID
	ExpiresAt  *string
	Scope      *string
}

// RevokeMemberRoleCmd 撤销成员角色命令
type RevokeMemberRoleCmd struct {
	MemberRoleID uuid.UUID
	RevokedBy    uuid.UUID
	RevokeReason string
}

// DeleteMemberRoleCmd 删除成员角色命令
type DeleteMemberRoleCmd struct {
	MemberRoleID uuid.UUID
}
