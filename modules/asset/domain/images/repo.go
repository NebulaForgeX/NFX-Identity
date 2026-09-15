package images

import (
	"context"

	"github.com/google/uuid"
)

type Repo struct {
	Create Create
	Get    Get
	Update Update
}

type Create interface {
	New(ctx context.Context, img *Image) error
}

type Get interface {
	ByID(ctx context.Context, id uuid.UUID) (*Image, error)
}

type Update interface {
	Generic(ctx context.Context, img *Image) error
}
