package phone

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
	New(ctx context.Context, p *Phone) error
}
type Get interface {
	ByID(ctx context.Context, id uuid.UUID) (*Phone, error)
	ByNumber(ctx context.Context, number string) (*Phone, error)
	ByAccountID(ctx context.Context, accountID uuid.UUID) ([]*Phone, error)
}
type Update interface {
	Generic(ctx context.Context, p *Phone) error
}
