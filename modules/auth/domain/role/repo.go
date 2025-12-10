package role

import (
	"context"

	"github.com/google/uuid"
)

type Repo interface {
	Create(ctx context.Context, r *Role) error
	Update(ctx context.Context, r *Role) error
	GetByID(ctx context.Context, id uuid.UUID) (*Role, error)
	GetByName(ctx context.Context, name string) (*Role, error)
	Exists(ctx context.Context, id uuid.UUID) (bool, error)
	ExistsByName(ctx context.Context, name string) (bool, error)
	Delete(ctx context.Context, id uuid.UUID) error
}

