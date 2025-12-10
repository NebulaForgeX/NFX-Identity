package image

import (
	"context"

	"github.com/google/uuid"
)

type Repo interface {
	Create(ctx context.Context, img *Image) error
	Update(ctx context.Context, img *Image) error
	GetByID(ctx context.Context, id uuid.UUID) (*Image, error)
	GetByFilename(ctx context.Context, filename string) (*Image, error)
	GetByStoragePath(ctx context.Context, storagePath string) (*Image, error)
	Exists(ctx context.Context, id uuid.UUID) (bool, error)
	Delete(ctx context.Context, id uuid.UUID) error
	DeleteByStoragePath(ctx context.Context, storagePath string) error
}

