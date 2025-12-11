package user_permission

import (
	"context"
	userPermissionViews "nfxid/modules/permission/application/user_permission/views"
	"github.com/google/uuid"
)

type UserPermissionQuery interface {
	GetByUserID(ctx context.Context, userID uuid.UUID) ([]*userPermissionViews.UserPermissionView, error)
	GetPermissionTagsByUserID(ctx context.Context, userID uuid.UUID) ([]string, error)
}

