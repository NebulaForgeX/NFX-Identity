package profile

import (
	"context"

	"nfxidentity/enums"
	"nfxidentity/pkgs/httpx"

	"github.com/google/uuid"
)

type AuthorityList interface {
	ByAccountID(ctx context.Context, accountID uuid.UUID) ([]AuthorityProfileItemVO, error)
	Search(ctx context.Context, q ListQuery) (httpx.Page[AuthorityProfileItemVO], error)
	BatchGetByProfileIDs(ctx context.Context, profileIDs []uuid.UUID) ([]AuthorityProfileItemVO, error)
	ProfileIDsByAuthorityRoles(ctx context.Context, roles []enums.AuthAuthorityRole) ([]uuid.UUID, error)
}

type ForgerList interface {
	ByAccountID(ctx context.Context, accountID uuid.UUID) ([]ForgerProfileItemVO, error)
	Search(ctx context.Context, q ListQuery) (httpx.Page[ForgerProfileItemVO], error)
	BatchGetByProfileIDs(ctx context.Context, profileIDs []uuid.UUID) ([]ForgerProfileItemVO, error)
}
