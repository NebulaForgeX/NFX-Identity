package account

import (
	"context"

	"nfxidentity/constants"
	"nfxidentity/enums"
	"nfxidentity/errors/src/auth"
	accountinternal "nfxidentity/modules/auth/application/account/internal"

	"github.com/google/uuid"
)

func (s *Service) EnsureOwnedProfile(ctx context.Context, accountID, profileID uuid.UUID, profileScope enums.AuthProfileScope) (bool, error) {
	if !constants.AuthProfileScope.Valid(profileScope) {
		return false, auth.ErrProfileScopeInvalid.WithDetail("profile_scope", string(profileScope))
	}
	switch profileScope {
	case enums.AuthProfileScopeCommunity:
		p, err := s.repoFactory.Profile(accountinternal.None()).Get.ForgerByProfileID(ctx, profileID)
		if err != nil {
			return false, err
		}
		return p.AccountID() == accountID, nil
	case enums.AuthProfileScopeAuthority:
		p, err := s.repoFactory.Profile(accountinternal.None()).Get.AuthorityByProfileID(ctx, profileID)
		if err != nil {
			return false, err
		}
		return p.AccountID() == accountID, nil
	default:
		return false, auth.ErrProfileScopeInvalid.WithDetail("profile_scope", string(profileScope))
	}
}

func (s *Service) requireOwnedProfile(ctx context.Context, accountID, profileID uuid.UUID, profileScope enums.AuthProfileScope) error {
	owned, err := s.EnsureOwnedProfile(ctx, accountID, profileID, profileScope)
	if err != nil {
		return err
	}
	if !owned {
		return auth.ErrProfileNotOwned
	}
	return nil
}
