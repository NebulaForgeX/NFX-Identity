package email

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
	New(ctx context.Context, e *Email) error
}

type Get interface {
	ByID(ctx context.Context, id uuid.UUID) (*Email, error)
	ByEmail(ctx context.Context, normalizedEmail string) (*Email, error)
	ByAccountID(ctx context.Context, accountID uuid.UUID) ([]*Email, error)
}

type Check interface {
	ByID(ctx context.Context, id uuid.UUID) (bool, error)
	ByEmail(ctx context.Context, normalizedEmail string) (bool, error)
	CountByAccountID(ctx context.Context, accountID uuid.UUID) (int64, error)
	CountVerifiedByAccountID(ctx context.Context, accountID uuid.UUID) (int64, error)
}

type Update interface {
	Generic(ctx context.Context, e *Email) error
}

type Delete interface {
	ByID(ctx context.Context, id uuid.UUID) error
}
