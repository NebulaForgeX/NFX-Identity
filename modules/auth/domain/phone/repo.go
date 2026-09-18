package phone

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
	New(ctx context.Context, p *Phone) error
}

type Get interface {
	ByID(ctx context.Context, id uuid.UUID) (*Phone, error)
	ByPhone(ctx context.Context, canonicalPhone string) (*Phone, error)
	ByAccountID(ctx context.Context, accountID uuid.UUID) ([]*Phone, error)
}

type Check interface {
	ByID(ctx context.Context, id uuid.UUID) (bool, error)
}

type Update interface {
	Generic(ctx context.Context, p *Phone) error
}

type Delete interface {
	ByID(ctx context.Context, id uuid.UUID) error
}
