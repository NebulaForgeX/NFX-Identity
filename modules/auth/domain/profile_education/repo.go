package education

import (
	"context"

	"github.com/google/uuid"
)

type Repo interface {
	Create(ctx context.Context, e *Education) error
	Update(ctx context.Context, e *Education) error
	GetByID(ctx context.Context, id uuid.UUID) (*Education, error)
	GetByProfileID(ctx context.Context, profileID uuid.UUID) ([]*Education, error)
	Exists(ctx context.Context, id uuid.UUID) (bool, error)
	Delete(ctx context.Context, id uuid.UUID) error
}

