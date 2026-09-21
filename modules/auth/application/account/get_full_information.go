package account

import (
	"context"

	accountQuery "nfxidentity/modules/auth/query/account"

	"github.com/google/uuid"
)

type GetFullInformationWithAuthorityProfileInput struct {
	AccountID uuid.UUID
	ProfileID uuid.UUID
}

type GetFullInformationWithCommunityProfileInput struct {
	AccountID uuid.UUID
	ProfileID uuid.UUID
}

func (s *Service) GetFullInformationWithAuthorityProfile(
	ctx context.Context,
	in GetFullInformationWithAuthorityProfileInput,
) (*accountQuery.FullAccountInformationWithAuthorityProfileVO, error) {
	return s.accounts.Single.ByIDWithAuthorityProfile(ctx, in.AccountID, in.ProfileID)
}

func (s *Service) GetFullInformationWithCommunityProfile(
	ctx context.Context,
	in GetFullInformationWithCommunityProfileInput,
) (*accountQuery.FullAccountInformationWithCommunityProfileVO, error) {
	return s.accounts.Single.ByIDWithCommunityProfile(ctx, in.AccountID, in.ProfileID)
}
