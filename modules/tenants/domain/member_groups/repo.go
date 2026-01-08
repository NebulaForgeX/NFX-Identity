package member_groups

import (
	"context"

	"github.com/google/uuid"
)

// Repo 是 MemberGroup 的仓库结构体，包含增删改查五个子接口
type Repo struct {
	Create Create
	Get    Get
	Check  Check
	Update Update
	Delete Delete
}

// Create 定义创建相关的方法
type Create interface {
	New(ctx context.Context, mg *MemberGroup) error
}

// Get 定义获取数据相关的方法
type Get interface {
	ByID(ctx context.Context, id uuid.UUID) (*MemberGroup, error)
	ByMemberID(ctx context.Context, memberID uuid.UUID) ([]*MemberGroup, error)
	ByGroupID(ctx context.Context, groupID uuid.UUID) ([]*MemberGroup, error)
	ByMemberIDAndGroupID(ctx context.Context, memberID, groupID uuid.UUID) (*MemberGroup, error)
	ActiveByMemberID(ctx context.Context, memberID uuid.UUID) ([]*MemberGroup, error)
}

// Check 定义检查相关的方法
type Check interface {
	ByID(ctx context.Context, id uuid.UUID) (bool, error)
	ByMemberIDAndGroupID(ctx context.Context, memberID, groupID uuid.UUID) (bool, error)
}

// Update 定义更新相关的方法
type Update interface {
	Generic(ctx context.Context, mg *MemberGroup) error
	Revoke(ctx context.Context, id uuid.UUID, revokedBy uuid.UUID) error
}

// Delete 定义删除相关的方法
type Delete interface {
	ByID(ctx context.Context, id uuid.UUID) error
	ByMemberIDAndGroupID(ctx context.Context, memberID, groupID uuid.UUID) error
}
