package domain_verifications

import (
	"context"

	"github.com/google/uuid"
)

// Repo 是 DomainVerification 的仓库结构体，包含增删改查五个子接口
type Repo struct {
	Create Create
	Get    Get
	Check  Check
	Update Update
	Delete Delete
}

// Create 定义创建相关的方法
type Create interface {
	New(ctx context.Context, dv *DomainVerification) error
}

// Get 定义获取数据相关的方法
type Get interface {
	ByID(ctx context.Context, id uuid.UUID) (*DomainVerification, error)
	ByTenantID(ctx context.Context, tenantID uuid.UUID) ([]*DomainVerification, error)
	ByDomain(ctx context.Context, domain string) (*DomainVerification, error)
	ByTenantIDAndDomain(ctx context.Context, tenantID uuid.UUID, domain string) (*DomainVerification, error)
	ByStatus(ctx context.Context, status VerificationStatus) ([]*DomainVerification, error)
}

// Check 定义检查相关的方法
type Check interface {
	ByID(ctx context.Context, id uuid.UUID) (bool, error)
	ByTenantIDAndDomain(ctx context.Context, tenantID uuid.UUID, domain string) (bool, error)
}

// Update 定义更新相关的方法
type Update interface {
	Generic(ctx context.Context, dv *DomainVerification) error
	Verify(ctx context.Context, id uuid.UUID) error
	Fail(ctx context.Context, id uuid.UUID) error
}

// Delete 定义删除相关的方法
type Delete interface {
	ByID(ctx context.Context, id uuid.UUID) error
}
