package profile

import (
	"context"

	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type Repo struct {
	Create Create
	Get    Get
	Check  Check
	Update Update
	Delete Delete
}

type Create interface {
	NewAuthority(ctx context.Context, p *AuthorityProfile) error
	NewForger(ctx context.Context, p *ForgerProfile) error
	NewAuthorityAvatar(ctx context.Context, a *AuthorityProfileAvatar) error
	NewForgerAvatar(ctx context.Context, a *ForgerProfileAvatar) error
	NewAuthorityBackground(ctx context.Context, b *AuthorityProfileBackground) error
	NewForgerBackground(ctx context.Context, b *ForgerProfileBackground) error
	NewAuthoritySettings(ctx context.Context, s *AuthorityProfileSettings) error
	NewForgerSettings(ctx context.Context, s *ForgerProfileSettings) error
}

type Get interface {
	AuthorityByProfileID(ctx context.Context, profileID uuid.UUID) (*AuthorityProfile, error)
	ForgerByProfileID(ctx context.Context, profileID uuid.UUID) (*ForgerProfile, error)
	AuthorityByAccountID(ctx context.Context, accountID uuid.UUID) (*AuthorityProfile, error)
	ForgerByAccountID(ctx context.Context, accountID uuid.UUID) (*ForgerProfile, error)
	AuthorityAvatarByProfileID(ctx context.Context, profileID uuid.UUID) (*AuthorityProfileAvatar, error)
	ForgerAvatarByProfileID(ctx context.Context, profileID uuid.UUID) (*ForgerProfileAvatar, error)
	ListAuthorityAvatarsByProfileID(ctx context.Context, profileID uuid.UUID) ([]*AuthorityProfileAvatar, error)
	ListForgerAvatarsByProfileID(ctx context.Context, profileID uuid.UUID) ([]*ForgerProfileAvatar, error)
	AuthorityBackgroundsByProfileID(ctx context.Context, profileID uuid.UUID) ([]*AuthorityProfileBackground, error)
	ForgerBackgroundsByProfileID(ctx context.Context, profileID uuid.UUID) ([]*ForgerProfileBackground, error)
	AuthorityBackgroundByRowID(ctx context.Context, rowID uuid.UUID) (*AuthorityProfileBackground, error)
	ForgerBackgroundByRowID(ctx context.Context, rowID uuid.UUID) (*ForgerProfileBackground, error)
	AuthoritySettingsByProfileID(ctx context.Context, profileID uuid.UUID) (*AuthorityProfileSettings, error)
	ForgerSettingsByProfileID(ctx context.Context, profileID uuid.UUID) (*ForgerProfileSettings, error)
	AuthorityByAccountAndID(ctx context.Context, accountID, profileID uuid.UUID) (*AuthorityProfile, error)
	ForgerByAccountAndID(ctx context.Context, accountID, profileID uuid.UUID) (*ForgerProfile, error)
	HasOwnerRole(ctx context.Context, accountID uuid.UUID) (bool, error)
	AnyOwnerExists(ctx context.Context) (bool, error)
}

type Check interface {
	AuthorityByAccountID(ctx context.Context, accountID uuid.UUID) (bool, error)
	ForgerByAccountID(ctx context.Context, accountID uuid.UUID) (bool, error)
	CountAuthorityByAccountID(ctx context.Context, accountID uuid.UUID) (int64, error)
	CountForgerByAccountID(ctx context.Context, accountID uuid.UUID) (int64, error)
}

type Update interface {
	AuthorityGeneric(ctx context.Context, p *AuthorityProfile) error
	ForgerGeneric(ctx context.Context, p *ForgerProfile) error
	AuthorityPartial(ctx context.Context, profileID uuid.UUID, p AuthorityProfilePatch) error
	ForgerPartial(ctx context.Context, profileID uuid.UUID, p ForgerProfilePatch) error
	AuthorityPreference(ctx context.Context, profileID uuid.UUID, preference *datatypes.JSON) error
	ForgerPreference(ctx context.Context, profileID uuid.UUID, preference *datatypes.JSON) error
	AuthorityAvatarGeneric(ctx context.Context, a *AuthorityProfileAvatar) error
	ForgerAvatarGeneric(ctx context.Context, a *ForgerProfileAvatar) error
	AuthorityBackgroundGeneric(ctx context.Context, b *AuthorityProfileBackground) error
	ForgerBackgroundGeneric(ctx context.Context, b *ForgerProfileBackground) error
	AuthorityPartialSettings(ctx context.Context, profileID uuid.UUID, p AuthorityProfileSettingsPatch) error
	ForgerPartialSettings(ctx context.Context, profileID uuid.UUID, p ForgerProfileSettingsPatch) error
}

type Delete interface {
	AuthorityByProfileID(ctx context.Context, profileID uuid.UUID) error
	ForgerByProfileID(ctx context.Context, profileID uuid.UUID) error
	DeleteAuthorityBackgroundsByProfileID(ctx context.Context, profileID uuid.UUID) error
	DeleteForgerBackgroundsByProfileID(ctx context.Context, profileID uuid.UUID) error
	DeleteAuthorityAvatarsByProfileID(ctx context.Context, profileID uuid.UUID) error
	DeleteForgerAvatarsByProfileID(ctx context.Context, profileID uuid.UUID) error
}
