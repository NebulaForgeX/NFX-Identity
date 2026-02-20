package application_roles

import (
	"context"

	"github.com/google/uuid"
)

type Repo struct {
	Create Create
	Get    Get
	Check  Check
	Update Update
	Delete Delete
}

type Create interface {
	New(ctx context.Context, r *ApplicationRole) error
}

type Get interface {
	ByID(ctx context.Context, id uuid.UUID) (*ApplicationRole, error)
	ByApplicationIDAndRoleKey(ctx context.Context, applicationID uuid.UUID, roleKey string) (*ApplicationRole, error)
	ByApplicationID(ctx context.Context, applicationID uuid.UUID) ([]*ApplicationRole, error)
}

type Check interface {
	ByID(ctx context.Context, id uuid.UUID) (bool, error)
	ByApplicationIDAndRoleKey(ctx context.Context, applicationID uuid.UUID, roleKey string) (bool, error)
}

type Update interface {
	Generic(ctx context.Context, r *ApplicationRole) error
}

type Delete interface {
	ByID(ctx context.Context, id uuid.UUID) error
}
