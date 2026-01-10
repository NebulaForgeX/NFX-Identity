package commands

import (
	"github.com/google/uuid"
)

// CreateMemberAppRoleCmd 创建成员应用角色命令
type CreateMemberAppRoleCmd struct {
	MemberID   uuid.UUID
	AppID      uuid.UUID
	RoleID     uuid.UUID
	AssignedBy *uuid.UUID
	ExpiresAt  *string
}

// RevokeMemberAppRoleCmd 撤销成员应用角色命令
type RevokeMemberAppRoleCmd struct {
	MemberAppRoleID uuid.UUID
	RevokedBy       uuid.UUID
	RevokeReason    string
}

// DeleteMemberAppRoleCmd 删除成员应用角色命令
type DeleteMemberAppRoleCmd struct {
	MemberAppRoleID uuid.UUID
}
