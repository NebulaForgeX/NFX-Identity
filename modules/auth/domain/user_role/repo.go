package user_role

import (
	"context"

	"github.com/google/uuid"
)

type Repo interface {
	Create(ctx context.Context, ur *UserRole) error
	GetByID(ctx context.Context, id uuid.UUID) (*UserRole, error)
	GetByUserID(ctx context.Context, userID uuid.UUID) ([]*UserRole, error)
	GetByRoleID(ctx context.Context, roleID uuid.UUID) ([]*UserRole, error)
	Exists(ctx context.Context, userID, roleID uuid.UUID) (bool, error)
	Delete(ctx context.Context, id uuid.UUID) error
	DeleteByUserAndRole(ctx context.Context, userID, roleID uuid.UUID) error
}

