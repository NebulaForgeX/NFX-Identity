package badge

import (
	"context"

	"github.com/google/uuid"
)

type Repo interface {
	Create(ctx context.Context, b *Badge) error
	Update(ctx context.Context, b *Badge) error
	GetByID(ctx context.Context, id uuid.UUID) (*Badge, error)
}
