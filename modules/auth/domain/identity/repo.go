package identity

import (
	"context"

	"nfxidentity/enums"

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
	New(ctx context.Context, i *Identity) error
}

type Get interface {
	ByID(ctx context.Context, id uuid.UUID) (*Identity, error)
	ByAccountID(ctx context.Context, accountID uuid.UUID) ([]*Identity, error)
	ByProviderSubject(ctx context.Context, provider enums.AuthIdentityProvider, subject string) (*Identity, error)
}

type Check interface {
	ByID(ctx context.Context, id uuid.UUID) (bool, error)
}

type Update interface {
	Generic(ctx context.Context, i *Identity) error
}

type Delete interface {
	ByID(ctx context.Context, id uuid.UUID) error
}
