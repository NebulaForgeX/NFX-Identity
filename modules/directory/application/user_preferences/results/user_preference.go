package results

import (
	"time"

	"nfxid/modules/directory/domain/user_preferences"

	"github.com/google/uuid"
)

type UserPreferenceRO struct {
	ID            uuid.UUID
	UserID        uuid.UUID
	Theme         string
	Language      string
	Timezone      string
	Notifications map[string]interface{}
	Privacy       map[string]interface{}
	Display       map[string]interface{}
	Other         map[string]interface{}
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     *time.Time
}

// UserPreferenceMapper 将 Domain UserPreference 转换为 Application UserPreferenceRO
func UserPreferenceMapper(up *user_preferences.UserPreference) UserPreferenceRO {
	if up == nil {
		return UserPreferenceRO{}
	}

	return UserPreferenceRO{
		ID:            up.ID(),
		UserID:        up.UserID(),
		Theme:         up.Theme(),
		Language:      up.Language(),
		Timezone:      up.Timezone(),
		Notifications: up.Notifications(),
		Privacy:       up.Privacy(),
		Display:       up.Display(),
		Other:         up.Other(),
		CreatedAt:     up.CreatedAt(),
		UpdatedAt:     up.UpdatedAt(),
		DeletedAt:     up.DeletedAt(),
	}
}
