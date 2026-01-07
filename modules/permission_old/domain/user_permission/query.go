package user_permission

import (
	"context"
	"nfxid/modules/permission/domain/user_permission/views"

	"github.com/google/uuid"
)

// Query 定义查询领域视图的结构体（CQRS Read Side）
type Query struct {
	Single Single
	List   List
}

// Single 定义单个查询相关的方法
type Single interface {
	ByUserIDAndPermissionID(ctx context.Context, userID, permissionID uuid.UUID) (*views.UserPermissionView, error)
}

// List 定义列表查询相关的方法
type List interface {
	ByUserID(ctx context.Context, userID uuid.UUID) ([]*views.UserPermissionView, error)
	PermissionTagsByUserID(ctx context.Context, userID uuid.UUID) ([]string, error)
}
