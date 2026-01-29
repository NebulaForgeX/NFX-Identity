package actions

import (
	"context"

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
	New(ctx context.Context, a *Action) error
}

type Get interface {
	ByID(ctx context.Context, id uuid.UUID) (*Action, error)
	ByKey(ctx context.Context, key string) (*Action, error)
}

type Check interface {
	ByID(ctx context.Context, id uuid.UUID) (bool, error)
	ByKey(ctx context.Context, key string) (bool, error)
}

type Update interface {
	Generic(ctx context.Context, a *Action) error
}

type Delete interface {
	ByID(ctx context.Context, id uuid.UUID) error
}
