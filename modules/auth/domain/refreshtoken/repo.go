package refreshtoken

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type Repo struct {
	Create Create
	Get    Get
	Update Update
}

type Create interface {
	New(ctx context.Context, t *RefreshToken) error
}
type Get interface {
	ByTokenHash(ctx context.Context, hash string) (*RefreshToken, error)
	ActiveByAccountDevice(ctx context.Context, accountID uuid.UUID, deviceID string) ([]*RefreshToken, error)
}
type Update interface {
	Generic(ctx context.Context, t *RefreshToken) error
	RevokeDevice(ctx context.Context, accountID uuid.UUID, deviceID string, at time.Time) error
}
