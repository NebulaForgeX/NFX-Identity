package commands

import (
	"github.com/google/uuid"
)

// CreateMemberGroupCmd 创建成员组命令
type CreateMemberGroupCmd struct {
	MemberID   uuid.UUID
	GroupID    uuid.UUID
	AssignedBy *uuid.UUID
}

// RevokeMemberGroupCmd 撤销成员组命令
type RevokeMemberGroupCmd struct {
	MemberGroupID uuid.UUID
	RevokedBy     uuid.UUID
}

// DeleteMemberGroupCmd 删除成员组命令
type DeleteMemberGroupCmd struct {
	MemberGroupID uuid.UUID
}
