package member_app_roles

import (
	"context"

	"github.com/google/uuid"
)

// Repo 是 MemberAppRole 的仓库结构体，包含增删改查五个子接口
type Repo struct {
	Create Create
	Get    Get
	Check  Check
	Update Update
	Delete Delete
}

// Create 定义创建相关的方法
type Create interface {
	New(ctx context.Context, mar *MemberAppRole) error
}

// Get 定义获取数据相关的方法
type Get interface {
	ByID(ctx context.Context, id uuid.UUID) (*MemberAppRole, error)
	ByMemberID(ctx context.Context, memberID uuid.UUID) ([]*MemberAppRole, error)
	ByAppID(ctx context.Context, appID uuid.UUID) ([]*MemberAppRole, error)
	ByMemberIDAndAppID(ctx context.Context, memberID, appID uuid.UUID) ([]*MemberAppRole, error)
	ByMemberIDAndAppIDAndRoleID(ctx context.Context, memberID, appID, roleID uuid.UUID) (*MemberAppRole, error)
	ActiveByMemberID(ctx context.Context, memberID uuid.UUID) ([]*MemberAppRole, error)
}

// Check 定义检查相关的方法
type Check interface {
	ByID(ctx context.Context, id uuid.UUID) (bool, error)
	ByMemberIDAndAppIDAndRoleID(ctx context.Context, memberID, appID, roleID uuid.UUID) (bool, error)
}

// Update 定义更新相关的方法
type Update interface {
	Generic(ctx context.Context, mar *MemberAppRole) error
	Revoke(ctx context.Context, id uuid.UUID, revokedBy uuid.UUID, reason string) error
}

// Delete 定义删除相关的方法
type Delete interface {
	ByID(ctx context.Context, id uuid.UUID) error
	ByMemberIDAndAppIDAndRoleID(ctx context.Context, memberID, appID, roleID uuid.UUID) error
}
