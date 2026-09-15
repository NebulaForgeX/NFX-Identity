package platform

import (
	"time"

	"nfxidentity/modules/auth/domain/avatar"
	"nfxidentity/modules/auth/domain/background"
	"nfxidentity/modules/auth/domain/refreshtoken"

	"github.com/google/uuid"
)

func refreshtokenNew(accountID uuid.UUID, identityID, profileID *uuid.UUID, scope, deviceID *string, hash string, now time.Time) *refreshtoken.RefreshToken {
	return refreshtoken.NewFromState(refreshtoken.RefreshTokenState{
		ID: uuid.New(), AccountID: accountID, IdentityID: identityID, ProfileID: profileID,
		ProfileScope: scope, DeviceID: deviceID, TokenHash: hash, ExpiresAt: now.Add(30 * 24 * time.Hour), CreatedAt: now,
	})
}

func newAvatar(profileID, imageID uuid.UUID, kind string, now time.Time) *avatar.Avatar {
	return avatar.NewFromState(avatar.State{
		ID: uuid.New(), ProfileID: profileID, ImageID: imageID, IsActive: true, Kind: kind, CreatedAt: now, UpdatedAt: now,
	})
}

func newBackground(profileID, imageID uuid.UUID, kind string, sort int, now time.Time) *background.Background {
	return background.NewFromState(background.State{
		ID: uuid.New(), ProfileID: profileID, ImageID: imageID, SortOrder: sort, Kind: kind, CreatedAt: now, UpdatedAt: now,
	})
}
