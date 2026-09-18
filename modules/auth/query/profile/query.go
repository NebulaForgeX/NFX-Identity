package profile

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

type ForgerItemVO struct {
	ProfileID       uuid.UUID      `json:"profile_id"`
	AccountID       uuid.UUID      `json:"account_id"`
	ForgerRoles     pq.StringArray `json:"forger_roles"`
	DisplayName     *string        `json:"display_name"`
	ProfileLanguage string         `json:"profile_language"`
	City            *string        `json:"city"`
	Country         *string        `json:"country"`
	Website         *string        `json:"website"`
	Timezone        *string        `json:"timezone"`
	AvatarImageID   *uuid.UUID     `json:"avatar_image_id"`
	CreatedAt       time.Time      `json:"created_at"`
}

type AuthorityItemVO struct {
	ProfileID       uuid.UUID      `json:"profile_id"`
	AccountID       uuid.UUID      `json:"account_id"`
	AuthorityRoles  pq.StringArray `json:"authority_roles"`
	DisplayName     *string        `json:"display_name"`
	ProfileLanguage string         `json:"profile_language"`
	City            *string        `json:"city"`
	Country         *string        `json:"country"`
	AvatarImageID   *uuid.UUID     `json:"avatar_image_id"`
	CreatedAt       time.Time      `json:"created_at"`
}

type Query struct {
	Forger    Forger
	Authority Authority
}

type Forger interface {
	ByAccountID(ctx context.Context, accountID uuid.UUID) ([]ForgerItemVO, error)
	Search(ctx context.Context, q string, limit, offset int) ([]ForgerItemVO, int64, error)
}

type Authority interface {
	ByAccountID(ctx context.Context, accountID uuid.UUID) ([]AuthorityItemVO, error)
	Search(ctx context.Context, q string, limit, offset int) ([]AuthorityItemVO, int64, error)
}
