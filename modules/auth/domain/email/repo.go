package email

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
	New(ctx context.Context, e *Email) error
}

type Get interface {
	ByID(ctx context.Context, id uuid.UUID) (*Email, error)
	ByAddress(ctx context.Context, address string) (*Email, error)
	ByAccountID(ctx context.Context, accountID uuid.UUID) ([]*Email, error)
}

type Update interface {
	Generic(ctx context.Context, e *Email) error
}
