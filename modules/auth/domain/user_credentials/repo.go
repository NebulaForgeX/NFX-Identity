package user_credentials

import (
	"context"

	"github.com/google/uuid"
)

// Repo 是 UserCredential 的仓库结构体，包含增删改查五个子接口
type Repo struct {
	Create Create
	Get    Get
	Check  Check
	Update Update
	Delete Delete
}

// Create 定义创建相关的方法
type Create interface {
	New(ctx context.Context, uc *UserCredential) error
}

// Get 定义获取数据相关的方法
type Get interface {
	ByID(ctx context.Context, id uuid.UUID) (*UserCredential, error)
	ByUserID(ctx context.Context, userID uuid.UUID) (*UserCredential, error)
}

// Check 定义检查相关的方法
type Check interface {
	ByID(ctx context.Context, id uuid.UUID) (bool, error)
	ByUserID(ctx context.Context, userID uuid.UUID) (bool, error)
}

// Update 定义更新相关的方法
type Update interface {
	Generic(ctx context.Context, uc *UserCredential) error
	UpdatePassword(ctx context.Context, userID uuid.UUID, passwordHash string, hashAlg string, hashParams map[string]interface{}) error
	UpdateStatus(ctx context.Context, userID uuid.UUID, status CredentialStatus) error
	UpdateLastSuccessLogin(ctx context.Context, userID uuid.UUID) error
}

// Delete 定义删除相关的方法
type Delete interface {
	ByID(ctx context.Context, id uuid.UUID) error
	ByUserID(ctx context.Context, userID uuid.UUID) error
}
