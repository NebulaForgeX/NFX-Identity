package profile_badge

import (
	"context"

	"github.com/google/uuid"
)

type Repo interface {
	Create(ctx context.Context, pb *ProfileBadge) error
	Update(ctx context.Context, pb *ProfileBadge) error
	GetByID(ctx context.Context, id uuid.UUID) (*ProfileBadge, error)
	GetByProfileID(ctx context.Context, profileID uuid.UUID) ([]*ProfileBadge, error)
	GetByBadgeID(ctx context.Context, badgeID uuid.UUID) ([]*ProfileBadge, error)
	Exists(ctx context.Context, profileID, badgeID uuid.UUID) (bool, error)
	Delete(ctx context.Context, id uuid.UUID) error
	DeleteByProfileAndBadge(ctx context.Context, profileID, badgeID uuid.UUID) error
}

