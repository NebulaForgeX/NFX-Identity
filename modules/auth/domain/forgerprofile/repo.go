package forgerprofile

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
	New(ctx context.Context, p *Profile) error
}
type Get interface {
	ByID(ctx context.Context, id uuid.UUID) (*Profile, error)
	ByAccountAndID(ctx context.Context, accountID, id uuid.UUID) (*Profile, error)
	ByAccountID(ctx context.Context, accountID uuid.UUID) ([]*Profile, error)
	Owned(ctx context.Context, accountID, id uuid.UUID) (bool, error)
	HasOwnerRole(ctx context.Context, accountID uuid.UUID) (bool, error)
}
type Update interface {
	Generic(ctx context.Context, p *Profile) error
}
