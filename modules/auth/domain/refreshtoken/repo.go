package refreshtoken

import (
	"context"
	"time"

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
	New(ctx context.Context, t *RefreshToken) error
}

type Get interface {
	ByID(ctx context.Context, id uuid.UUID) (*RefreshToken, error)
	ByTokenHash(ctx context.Context, tokenHash string) (*RefreshToken, error)
	ByAccountIDAndDeviceID(ctx context.Context, accountID uuid.UUID, deviceID string) ([]*RefreshToken, error)
}

type Check interface {
	ByID(ctx context.Context, id uuid.UUID) (bool, error)
}

type Update interface {
	Generic(ctx context.Context, t *RefreshToken) error
	RevokeDevice(ctx context.Context, accountID uuid.UUID, deviceID string, at time.Time) error
}

type Delete interface {
	ByID(ctx context.Context, id uuid.UUID) error
}
