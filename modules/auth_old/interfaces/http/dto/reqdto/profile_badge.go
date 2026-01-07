package reqdto

import (
	profileBadgeAppCommands "nfxid/modules/auth/application/profile_badge/commands"
	profileBadgeDomain "nfxid/modules/auth/domain/profile_badge"

	"github.com/google/uuid"
)

type ProfileBadgeCreateRequestDTO struct {
	ProfileID   uuid.UUID `json:"profile_id" validate:"required"`
	BadgeID     uuid.UUID `json:"badge_id" validate:"required"`
	Description *string   `json:"description,omitempty"`
	Level       *int      `json:"level,omitempty"`
}

type ProfileBadgeUpdateRequestDTO struct {
	ID          uuid.UUID `params:"id" validate:"required,uuid"`
	Description *string   `json:"description,omitempty"`
	Level       *int      `json:"level,omitempty"`
}

type ProfileBadgeByIDRequestDTO struct {
	ID uuid.UUID `params:"id" validate:"required,uuid"`
}

type ProfileBadgeByProfileIDRequestDTO struct {
	ProfileID uuid.UUID `params:"profile_id" validate:"required,uuid"`
}

type ProfileBadgeByBadgeIDRequestDTO struct {
	BadgeID uuid.UUID `params:"badge_id" validate:"required,uuid"`
}

type ProfileBadgeDeleteByProfileAndBadgeRequestDTO struct {
	ProfileID uuid.UUID `params:"profile_id" validate:"required,uuid"`
	BadgeID   uuid.UUID `params:"badge_id" validate:"required,uuid"`
}

func (r *ProfileBadgeCreateRequestDTO) ToCreateCmd() profileBadgeAppCommands.CreateProfileBadgeCmd {
	editable := profileBadgeDomain.ProfileBadgeEditable{}
	if r.Description != nil {
		editable.Description = r.Description
	}
	if r.Level != nil {
		editable.Level = r.Level
	}

	return profileBadgeAppCommands.CreateProfileBadgeCmd{
		ProfileID: r.ProfileID,
		BadgeID:   r.BadgeID,
		Editable:  editable,
	}
}

func (r *ProfileBadgeUpdateRequestDTO) ToUpdateCmd() profileBadgeAppCommands.UpdateProfileBadgeCmd {
	editable := profileBadgeDomain.ProfileBadgeEditable{}
	if r.Description != nil {
		editable.Description = r.Description
	}
	if r.Level != nil {
		editable.Level = r.Level
	}

	return profileBadgeAppCommands.UpdateProfileBadgeCmd{
		ProfileBadgeID: r.ID,
		Editable:       editable,
	}
}
