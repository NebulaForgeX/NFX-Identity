package account

import (
	"context"

	accountQuery "nfxidentity/modules/auth/query/account"

	"github.com/google/uuid"
)

func (s *Service) GetPrimaryEmailByProfileID(ctx context.Context, profileID uuid.UUID) (*accountQuery.PrimaryEmailVO, error) {
	vo, err := s.accounts.Single.PrimaryEmailByCommunityProfileID(ctx, profileID)
	if err == nil {
		return vo, nil
	}
	return s.accounts.Single.PrimaryEmailByAuthorityProfileID(ctx, profileID)
}
