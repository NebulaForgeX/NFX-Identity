package image_type

import (
	"context"

	"github.com/google/uuid"
)

type Repo interface {
	Create(ctx context.Context, it *ImageType) error
	Update(ctx context.Context, it *ImageType) error
	GetByID(ctx context.Context, id uuid.UUID) (*ImageType, error)
	GetByKey(ctx context.Context, key string) (*ImageType, error)
	Exists(ctx context.Context, id uuid.UUID) (bool, error)
	ExistsByKey(ctx context.Context, key string) (bool, error)
	Delete(ctx context.Context, id uuid.UUID) error
}

