package identity

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
	New(ctx context.Context, i *Identity) error
}

type Get interface {
	ByID(ctx context.Context, id uuid.UUID) (*Identity, error)
	ByProviderSubject(ctx context.Context, provider, subject string) (*Identity, error)
	ByAccountID(ctx context.Context, accountID uuid.UUID) ([]*Identity, error)
}

type Update interface {
	Generic(ctx context.Context, i *Identity) error
}
