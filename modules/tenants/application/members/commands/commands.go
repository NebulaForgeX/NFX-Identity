package commands

import (
	"nfxid/modules/tenants/domain/members"

	"github.com/google/uuid"
)

// CreateMemberCmd 创建成员命令
type CreateMemberCmd struct {
	TenantID   uuid.UUID
	UserID     uuid.UUID
	Status     members.MemberStatus
	Source     members.MemberSource
	CreatedBy  *uuid.UUID
	ExternalRef *string
	Metadata   map[string]interface{}
}

// UpdateMemberStatusCmd 更新成员状态命令
type UpdateMemberStatusCmd struct {
	MemberID uuid.UUID
	Status   members.MemberStatus
}

// JoinMemberCmd 成员加入命令
type JoinMemberCmd struct {
	MemberID uuid.UUID
}

// LeaveMemberCmd 成员离开命令
type LeaveMemberCmd struct {
	MemberID uuid.UUID
}

// SuspendMemberCmd 暂停成员命令
type SuspendMemberCmd struct {
	MemberID uuid.UUID
}

// DeleteMemberCmd 删除成员命令
type DeleteMemberCmd struct {
	MemberID uuid.UUID
}
