package permission

import (
	"context"
	"nfxid/enums"
	permissionViews "nfxid/modules/permission/application/permission/views"

	"github.com/google/uuid"
)

type PermissionQuery interface {
	GetByID(ctx context.Context, id uuid.UUID) (*permissionViews.PermissionView, error)
	GetByTag(ctx context.Context, tag string) (*permissionViews.PermissionView, error)
	GetByTags(ctx context.Context, tags []string) ([]*permissionViews.PermissionView, error)
	GetByCategory(ctx context.Context, category enums.PermissionCategory) ([]*permissionViews.PermissionView, error)
	List(ctx context.Context) ([]*permissionViews.PermissionView, error)
}
