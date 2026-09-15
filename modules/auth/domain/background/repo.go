package background

import (
	"context"

	"github.com/google/uuid"
)

type Repo struct {
	Create Create
	Get    Get
	Delete Delete
}

type Create interface {
	New(ctx context.Context, b *Background) error
}
type Get interface {
	ByProfileID(ctx context.Context, kind string, profileID uuid.UUID) ([]*Background, error)
}
type Delete interface {
	AllByProfileID(ctx context.Context, kind string, profileID uuid.UUID) error
}
