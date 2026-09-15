package account

import (
	"context"

	"github.com/google/uuid"
)

type Repo struct {
	Create Create
	Get    Get
	Check  Check
	Update Update
}

type Create interface {
	New(ctx context.Context, a *Account) error
}

type Get interface {
	ByID(ctx context.Context, id uuid.UUID) (*Account, error)
}

type Check interface {
	ByID(ctx context.Context, id uuid.UUID) (bool, error)
}

type Update interface {
	Generic(ctx context.Context, a *Account) error
}
