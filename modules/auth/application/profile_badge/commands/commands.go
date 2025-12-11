package commands

import (
	profileBadgeDomain "nfxid/modules/auth/domain/profile_badge"

	"github.com/google/uuid"
)

type CreateProfileBadgeCmd struct {
	ProfileID uuid.UUID
	BadgeID   uuid.UUID
	Editable  profileBadgeDomain.ProfileBadgeEditable
}

type UpdateProfileBadgeCmd struct {
	ProfileBadgeID uuid.UUID
	Editable       profileBadgeDomain.ProfileBadgeEditable
}
