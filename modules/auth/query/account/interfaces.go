package account

import (
	"context"

	"github.com/google/uuid"
)

type Single interface {
	ByIDWithCommunityProfile(ctx context.Context, accountID, profileID uuid.UUID) (*FullAccountInformationWithCommunityProfileVO, error)
	ByIDWithAuthorityProfile(ctx context.Context, accountID, profileID uuid.UUID) (*FullAccountInformationWithAuthorityProfileVO, error)
	PrimaryEmailByCommunityProfileID(ctx context.Context, profileID uuid.UUID) (*PrimaryEmailVO, error)
	PrimaryEmailByAuthorityProfileID(ctx context.Context, profileID uuid.UUID) (*PrimaryEmailVO, error)
}
