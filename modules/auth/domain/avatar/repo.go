package avatar

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
	New(ctx context.Context, a *Avatar) error
}
type Get interface {
	ByProfileID(ctx context.Context, kind string, profileID uuid.UUID) ([]*Avatar, error)
}
type Update interface {
	DeactivateActive(ctx context.Context, kind string, profileID uuid.UUID) error
}
