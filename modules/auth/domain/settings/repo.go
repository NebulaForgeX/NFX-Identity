package settings

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
	New(ctx context.Context, s *Settings) error
}
type Get interface {
	ByID(ctx context.Context, kind string, id uuid.UUID) (*Settings, error)
}
type Update interface {
	Generic(ctx context.Context, s *Settings) error
}
