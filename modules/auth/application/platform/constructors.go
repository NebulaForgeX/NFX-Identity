package platform

import (
	"time"

	"nfxidentity/enums"
	"nfxidentity/modules/auth/domain/profile"
	"nfxidentity/modules/auth/domain/refreshtoken"

	"github.com/google/uuid"
)

func refreshtokenNew(accountID uuid.UUID, identityID, profileID *uuid.UUID, scope *enums.AuthProfileScope, deviceID *string, hash string, now time.Time) *refreshtoken.RefreshToken {
	return refreshtoken.NewRefreshTokenFromState(refreshtoken.RefreshTokenState{
		ID: uuid.New(), AccountID: accountID, IdentityID: identityID, ProfileID: profileID,
		ProfileScope: scope, DeviceID: deviceID, TokenHash: hash, ExpiresAt: now.Add(30 * 24 * time.Hour), CreatedAt: now,
	})
}

func newForgerAvatar(profileID, imageID uuid.UUID, now time.Time) *profile.ForgerProfileAvatar {
	return profile.NewForgerProfileAvatarFromState(profile.ForgerProfileAvatarState{
		ID: uuid.New(), ProfileID: profileID, ImageID: imageID, IsActive: true, CreatedAt: now, UpdatedAt: now,
	})
}

func newAuthorityAvatar(profileID, imageID uuid.UUID, now time.Time) *profile.AuthorityProfileAvatar {
	return profile.NewAuthorityProfileAvatarFromState(profile.AuthorityProfileAvatarState{
		ID: uuid.New(), ProfileID: profileID, ImageID: imageID, IsActive: true, CreatedAt: now, UpdatedAt: now,
	})
}

func newForgerBackground(profileID, imageID uuid.UUID, sort int, now time.Time) *profile.ForgerProfileBackground {
	return profile.NewForgerProfileBackgroundFromState(profile.ForgerProfileBackgroundState{
		ID: uuid.New(), ProfileID: profileID, ImageID: imageID, SortOrder: sort, CreatedAt: now, UpdatedAt: now,
	})
}

func newAuthorityBackground(profileID, imageID uuid.UUID, sort int, now time.Time) *profile.AuthorityProfileBackground {
	return profile.NewAuthorityProfileBackgroundFromState(profile.AuthorityProfileBackgroundState{
		ID: uuid.New(), ProfileID: profileID, ImageID: imageID, SortOrder: sort, CreatedAt: now, UpdatedAt: now,
	})
}

func scopePtr(kind string) *enums.AuthProfileScope {
	if kind == "" {
		return nil
	}
	s := enums.AuthProfileScope(kind)
	return &s
}

func roleStrings[T ~string](roles []T) []string {
	out := make([]string, 0, len(roles))
	for _, r := range roles {
		out = append(out, string(r))
	}
	return out
}
