package files

import (
	"context"
	"time"

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
	New(ctx context.Context, e *File) error
}

type Get interface {
	ByID(ctx context.Context, id uuid.UUID) (*File, error)
	ListStaleTmpBefore(ctx context.Context, cutoff time.Time, limit int) ([]*File, error)
}

type Check interface {
	ByID(ctx context.Context, id uuid.UUID) (bool, error)
}

type Update interface {
	Generic(ctx context.Context, e *File) error
}

type Delete interface {
	ByID(ctx context.Context, id uuid.UUID) error
}
