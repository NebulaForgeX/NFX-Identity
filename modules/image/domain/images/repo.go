package images

import (
	"context"

	"github.com/google/uuid"
)

// Repo 是 Image 的仓库结构体，包含增删改查五个子接口
type Repo struct {
	Create Create
	Get    Get
	Check  Check
	Update Update
	Delete Delete
}

// Create 定义创建相关的方法
type Create interface {
	New(ctx context.Context, i *Image) error
}

// Get 定义获取数据相关的方法
type Get interface {
	ByID(ctx context.Context, id uuid.UUID) (*Image, error)
	ByUserID(ctx context.Context, userID uuid.UUID) ([]*Image, error)
	ByTenantID(ctx context.Context, tenantID uuid.UUID) ([]*Image, error)
	ByTypeID(ctx context.Context, typeID uuid.UUID) ([]*Image, error)
	BySourceDomain(ctx context.Context, sourceDomain string) ([]*Image, error)
	PublicByUserID(ctx context.Context, userID uuid.UUID) ([]*Image, error)
}

// Check 定义检查相关的方法
type Check interface {
	ByID(ctx context.Context, id uuid.UUID) (bool, error)
}

// Update 定义更新相关的方法
type Update interface {
	Generic(ctx context.Context, i *Image) error
	UpdateURL(ctx context.Context, id uuid.UUID, url string) error
	UpdatePublic(ctx context.Context, id uuid.UUID, isPublic bool) error
}

// Delete 定义删除相关的方法
type Delete interface {
	ByID(ctx context.Context, id uuid.UUID) error
}
