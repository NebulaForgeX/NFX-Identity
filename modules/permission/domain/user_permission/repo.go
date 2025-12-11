package user_permission

import (
	"context"

	"github.com/google/uuid"
)

type Repo interface {
	Create(ctx context.Context, up *UserPermission) error
	GetByUserID(ctx context.Context, userID uuid.UUID) ([]*UserPermission, error)
	GetByUserIDAndPermissionID(ctx context.Context, userID, permissionID uuid.UUID) (*UserPermission, error)
	GetPermissionTagsByUserID(ctx context.Context, userID uuid.UUID) ([]string, error)
	Exists(ctx context.Context, userID, permissionID uuid.UUID) (bool, error)
	Delete(ctx context.Context, id uuid.UUID) error
	DeleteByUserIDAndPermissionID(ctx context.Context, userID, permissionID uuid.UUID) error
	DeleteByUserID(ctx context.Context, userID uuid.UUID) error
}

