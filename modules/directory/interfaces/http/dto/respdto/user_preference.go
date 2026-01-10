package respdto

import (
	"time"

	userPreferenceAppResult "nfxid/modules/directory/application/user_preferences/results"

	"github.com/google/uuid"
)

type UserPreferenceDTO struct {
	ID            uuid.UUID              `json:"id"`
	UserID        uuid.UUID              `json:"user_id"`
	Theme         string                 `json:"theme"`
	Language      string                 `json:"language"`
	Timezone      string                 `json:"timezone"`
	Notifications map[string]interface{} `json:"notifications,omitempty"`
	Privacy       map[string]interface{} `json:"privacy,omitempty"`
	Display       map[string]interface{} `json:"display,omitempty"`
	Other         map[string]interface{} `json:"other,omitempty"`
	CreatedAt     time.Time              `json:"created_at"`
	UpdatedAt     time.Time              `json:"updated_at"`
	DeletedAt     *time.Time             `json:"deleted_at,omitempty"`
}

// UserPreferenceROToDTO converts application UserPreferenceRO to response DTO
func UserPreferenceROToDTO(v *userPreferenceAppResult.UserPreferenceRO) *UserPreferenceDTO {
	if v == nil {
		return nil
	}

	return &UserPreferenceDTO{
		ID:            v.ID,
		UserID:        v.UserID,
		Theme:         v.Theme,
		Language:      v.Language,
		Timezone:      v.Timezone,
		Notifications: v.Notifications,
		Privacy:       v.Privacy,
		Display:       v.Display,
		Other:         v.Other,
		CreatedAt:     v.CreatedAt,
		UpdatedAt:     v.UpdatedAt,
		DeletedAt:     v.DeletedAt,
	}
}

// UserPreferenceListROToDTO converts list of UserPreferenceRO to DTOs
func UserPreferenceListROToDTO(results []userPreferenceAppResult.UserPreferenceRO) []UserPreferenceDTO {
	dtos := make([]UserPreferenceDTO, len(results))
	for i, v := range results {
		if dto := UserPreferenceROToDTO(&v); dto != nil {
			dtos[i] = *dto
		}
	}
	return dtos
}
