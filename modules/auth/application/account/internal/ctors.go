package internal

import (
	"time"

	"nfxidentity/modules/auth/domain/profile"

	"github.com/google/uuid"
)

func NewForgerAvatar(profileID, imageID uuid.UUID, now time.Time) *profile.ForgerProfileAvatar {
	return profile.NewForgerProfileAvatarFromState(profile.ForgerProfileAvatarState{
		ID: uuid.New(), ProfileID: profileID, ImageID: imageID, IsActive: true, CreatedAt: now, UpdatedAt: now,
	})
}

func NewAuthorityAvatar(profileID, imageID uuid.UUID, now time.Time) *profile.AuthorityProfileAvatar {
	return profile.NewAuthorityProfileAvatarFromState(profile.AuthorityProfileAvatarState{
		ID: uuid.New(), ProfileID: profileID, ImageID: imageID, IsActive: true, CreatedAt: now, UpdatedAt: now,
	})
}

func NewForgerBackground(profileID, imageID uuid.UUID, sort int, now time.Time) *profile.ForgerProfileBackground {
	return profile.NewForgerProfileBackgroundFromState(profile.ForgerProfileBackgroundState{
		ID: uuid.New(), ProfileID: profileID, ImageID: imageID, SortOrder: sort, CreatedAt: now, UpdatedAt: now,
	})
}

func NewAuthorityBackground(profileID, imageID uuid.UUID, sort int, now time.Time) *profile.AuthorityProfileBackground {
	return profile.NewAuthorityProfileBackgroundFromState(profile.AuthorityProfileBackgroundState{
		ID: uuid.New(), ProfileID: profileID, ImageID: imageID, SortOrder: sort, CreatedAt: now, UpdatedAt: now,
	})
}

func LimitOr(n, d int) int {
	if n <= 0 {
		return d
	}
	return n
}
