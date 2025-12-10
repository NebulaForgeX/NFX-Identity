package occupation

import (
	"context"

	"github.com/google/uuid"
)

type Repo interface {
	Create(ctx context.Context, o *Occupation) error
	Update(ctx context.Context, o *Occupation) error
	GetByID(ctx context.Context, id uuid.UUID) (*Occupation, error)
	GetByProfileID(ctx context.Context, profileID uuid.UUID) ([]*Occupation, error)
	Exists(ctx context.Context, id uuid.UUID) (bool, error)
	Delete(ctx context.Context, id uuid.UUID) error
}

