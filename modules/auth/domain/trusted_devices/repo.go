package trusted_devices

import (
	"context"
	"time"

	"github.com/google/uuid"
)

// Repo 是 TrustedDevice 的仓库结构体，包含增删改查五个子接口
type Repo struct {
	Create Create
	Get    Get
	Check  Check
	Update Update
	Delete Delete
}

// Create 定义创建相关的方法
type Create interface {
	New(ctx context.Context, td *TrustedDevice) error
}

// Get 定义获取数据相关的方法
type Get interface {
	ByID(ctx context.Context, id uuid.UUID) (*TrustedDevice, error)
	ByDeviceID(ctx context.Context, deviceID string) (*TrustedDevice, error)
	ByUserIDAndTenantID(ctx context.Context, userID, tenantID uuid.UUID) ([]*TrustedDevice, error)
	ByUserIDAndDeviceIDAndTenantID(ctx context.Context, userID uuid.UUID, deviceID string, tenantID uuid.UUID) (*TrustedDevice, error)
}

// Check 定义检查相关的方法
type Check interface {
	ByID(ctx context.Context, id uuid.UUID) (bool, error)
	ByUserIDAndDeviceIDAndTenantID(ctx context.Context, userID uuid.UUID, deviceID string, tenantID uuid.UUID) (bool, error)
	IsTrusted(ctx context.Context, userID uuid.UUID, deviceID string, tenantID uuid.UUID) (bool, error)
}

// Update 定义更新相关的方法
type Update interface {
	Generic(ctx context.Context, td *TrustedDevice) error
	UpdateLastUsed(ctx context.Context, deviceID string) error
	UpdateTrustedUntil(ctx context.Context, deviceID string, trustedUntil time.Time) error
}

// Delete 定义删除相关的方法
type Delete interface {
	ByID(ctx context.Context, id uuid.UUID) error
	ByDeviceID(ctx context.Context, deviceID string) error
	ByUserIDAndDeviceIDAndTenantID(ctx context.Context, userID uuid.UUID, deviceID string, tenantID uuid.UUID) error
}
