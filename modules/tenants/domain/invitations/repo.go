package invitations

import (
	"context"

	"github.com/google/uuid"
)

// Repo 是 Invitation 的仓库结构体，包含增删改查五个子接口
type Repo struct {
	Create Create
	Get    Get
	Check  Check
	Update Update
	Delete Delete
}

// Create 定义创建相关的方法
type Create interface {
	New(ctx context.Context, i *Invitation) error
}

// Get 定义获取数据相关的方法
type Get interface {
	ByID(ctx context.Context, id uuid.UUID) (*Invitation, error)
	ByInviteID(ctx context.Context, inviteID string) (*Invitation, error)
	ByTenantID(ctx context.Context, tenantID uuid.UUID) ([]*Invitation, error)
	ByEmail(ctx context.Context, email string) ([]*Invitation, error)
	ByStatus(ctx context.Context, status InvitationStatus) ([]*Invitation, error)
	ByTenantIDAndStatus(ctx context.Context, tenantID uuid.UUID, status InvitationStatus) ([]*Invitation, error)
}

// Check 定义检查相关的方法
type Check interface {
	ByID(ctx context.Context, id uuid.UUID) (bool, error)
	ByInviteID(ctx context.Context, inviteID string) (bool, error)
}

// Update 定义更新相关的方法
type Update interface {
	Generic(ctx context.Context, i *Invitation) error
	Accept(ctx context.Context, inviteID string, userID uuid.UUID) error
	Revoke(ctx context.Context, inviteID string, revokedBy uuid.UUID, reason string) error
}

// Delete 定义删除相关的方法
type Delete interface {
	ByID(ctx context.Context, id uuid.UUID) error
	ByInviteID(ctx context.Context, inviteID string) error
}
