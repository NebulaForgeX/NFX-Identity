package member_roles

import (
	"context"

	"github.com/google/uuid"
)

// Repo 是 MemberRole 的仓库结构体，包含增删改查五个子接口
type Repo struct {
	Create Create
	Get    Get
	Check  Check
	Update Update
	Delete Delete
}

// Create 定义创建相关的方法
type Create interface {
	New(ctx context.Context, mr *MemberRole) error
}

// Get 定义获取数据相关的方法
type Get interface {
	ByID(ctx context.Context, id uuid.UUID) (*MemberRole, error)
	ByMemberID(ctx context.Context, memberID uuid.UUID) ([]*MemberRole, error)
	ByRoleID(ctx context.Context, roleID uuid.UUID) ([]*MemberRole, error)
	ByTenantID(ctx context.Context, tenantID uuid.UUID) ([]*MemberRole, error)
	ByTenantIDAndMemberID(ctx context.Context, tenantID, memberID uuid.UUID) ([]*MemberRole, error)
	ActiveByMemberID(ctx context.Context, memberID uuid.UUID) ([]*MemberRole, error)
}

// Check 定义检查相关的方法
type Check interface {
	ByID(ctx context.Context, id uuid.UUID) (bool, error)
	ByTenantIDAndMemberIDAndRoleID(ctx context.Context, tenantID, memberID, roleID uuid.UUID) (bool, error)
}

// Update 定义更新相关的方法
type Update interface {
	Generic(ctx context.Context, mr *MemberRole) error
	Revoke(ctx context.Context, id uuid.UUID, revokedBy uuid.UUID, reason string) error
}

// Delete 定义删除相关的方法
type Delete interface {
	ByID(ctx context.Context, id uuid.UUID) error
	ByTenantIDAndMemberIDAndRoleID(ctx context.Context, tenantID, memberID, roleID uuid.UUID) error
}
