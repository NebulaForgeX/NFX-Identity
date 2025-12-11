package permission

import (
	"context"

	"github.com/google/uuid"
)

type Repo interface {
	Create(ctx context.Context, p *Permission) error
	Update(ctx context.Context, p *Permission) error
	GetByID(ctx context.Context, id uuid.UUID) (*Permission, error)
	GetByTag(ctx context.Context, tag string) (*Permission, error)
	GetByTags(ctx context.Context, tags []string) ([]*Permission, error)
	GetByCategory(ctx context.Context, category string) ([]*Permission, error)
	Exists(ctx context.Context, id uuid.UUID) (bool, error)
	ExistsByTag(ctx context.Context, tag string) (bool, error)
	Delete(ctx context.Context, id uuid.UUID) error
}

