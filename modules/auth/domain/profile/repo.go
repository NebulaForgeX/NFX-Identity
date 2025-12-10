package profile

import (
	"context"

	"github.com/google/uuid"
)

type Repo interface {
	Create(ctx context.Context, p *Profile) error
	Update(ctx context.Context, p *Profile) error
	GetByID(ctx context.Context, id uuid.UUID) (*Profile, error)
	GetByUserID(ctx context.Context, userID uuid.UUID) (*Profile, error)
	Exists(ctx context.Context, id uuid.UUID) (bool, error)
	ExistsByUserID(ctx context.Context, userID uuid.UUID) (bool, error)
	Delete(ctx context.Context, id uuid.UUID) error
}

