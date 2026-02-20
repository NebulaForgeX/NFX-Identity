package super_admins

import (
	"context"

	"github.com/google/uuid"
)

type Repo struct {
	Create Create
	Get    Get
	Check  Check
	Delete Delete
}

type Create interface {
	New(ctx context.Context, s *SuperAdmin) error
}

type Get interface {
	ByUserID(ctx context.Context, userID uuid.UUID) (*SuperAdmin, error)
	All(ctx context.Context, limit, offset int) ([]*SuperAdmin, error)
}

type Check interface {
	ByUserID(ctx context.Context, userID uuid.UUID) (bool, error)
}

type Delete interface {
	ByUserID(ctx context.Context, userID uuid.UUID) error
}
