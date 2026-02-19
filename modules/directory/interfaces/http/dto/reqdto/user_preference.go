package reqdto

import (
	userPreferenceAppCommands "nfxid/modules/directory/application/user_preferences/commands"

	"github.com/google/uuid"
)

type UserPreferenceCreateRequestDTO struct {
	UserID        uuid.UUID              `json:"user_id" validate:"required,uuid"`
	Theme         string                 `json:"theme" validate:"required"`
	Language      string                 `json:"language" validate:"required"`
	Timezone      string                 `json:"timezone" validate:"required"`
	Notifications map[string]interface{} `json:"notifications,omitempty"`
	Privacy       map[string]interface{} `json:"privacy,omitempty"`
	Display       map[string]interface{} `json:"display,omitempty"`
	Other         map[string]interface{} `json:"other,omitempty"`
}

type UserPreferenceUpdateRequestDTO struct {
	ID            uuid.UUID              `uri:"id" validate:"required,uuid"`
	Theme         *string                `json:"theme,omitempty"`
	Language      *string                `json:"language,omitempty"`
	Timezone      *string                `json:"timezone,omitempty"`
	Notifications map[string]interface{} `json:"notifications,omitempty"`
	Privacy       map[string]interface{} `json:"privacy,omitempty"`
	Display       map[string]interface{} `json:"display,omitempty"`
	Other         map[string]interface{} `json:"other,omitempty"`
}

func (r *UserPreferenceCreateRequestDTO) ToCreateCmd() userPreferenceAppCommands.CreateUserPreferenceCmd {
	return userPreferenceAppCommands.CreateUserPreferenceCmd{
		UserID:        r.UserID,
		Theme:         r.Theme,
		Language:      r.Language,
		Timezone:      r.Timezone,
		Notifications: r.Notifications,
		Privacy:       r.Privacy,
		Display:       r.Display,
		Other:         r.Other,
	}
}

func (r *UserPreferenceUpdateRequestDTO) ToUpdateCmd() userPreferenceAppCommands.UpdateUserPreferenceCmd {
	return userPreferenceAppCommands.UpdateUserPreferenceCmd{
		UserPreferenceID: r.ID,
		Theme:            r.Theme,
		Language:         r.Language,
		Timezone:         r.Timezone,
		Notifications:    r.Notifications,
		Privacy:          r.Privacy,
		Display:          r.Display,
		Other:            r.Other,
	}
}
