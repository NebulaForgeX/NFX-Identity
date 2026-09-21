package account

import (
	"context"

	"github.com/google/uuid"
	"gorm.io/datatypes"

	"nfxidentity/enums"
	accountinternal "nfxidentity/modules/auth/application/account/internal"
)

type UpdateCommunityPreferenceInput struct {
	AccountID  uuid.UUID
	ProfileID  uuid.UUID
	Preference string
}

func (s *Service) UpdateCommunityPreference(ctx context.Context, in UpdateCommunityPreferenceInput) error {
	if err := s.requireOwnedProfile(ctx, in.AccountID, in.ProfileID, enums.AuthProfileScopeCommunity); err != nil {
		return err
	}
	raw := datatypes.JSON([]byte(in.Preference))
	if in.Preference == "" {
		raw = datatypes.JSON([]byte("{}"))
	}
	return s.repoFactory.Profile(accountinternal.None()).Update.ForgerPreference(ctx, in.ProfileID, &raw)
}

type UpdateAuthorityPreferenceInput struct {
	AccountID  uuid.UUID
	ProfileID  uuid.UUID
	Preference string
}

func (s *Service) UpdateAuthorityPreference(ctx context.Context, in UpdateAuthorityPreferenceInput) error {
	if err := s.requireOwnedProfile(ctx, in.AccountID, in.ProfileID, enums.AuthProfileScopeAuthority); err != nil {
		return err
	}
	raw := datatypes.JSON([]byte(in.Preference))
	if in.Preference == "" {
		raw = datatypes.JSON([]byte("{}"))
	}
	return s.repoFactory.Profile(accountinternal.None()).Update.AuthorityPreference(ctx, in.ProfileID, &raw)
}
