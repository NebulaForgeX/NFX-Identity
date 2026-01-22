package user_preferences

import (
	"time"

	"github.com/google/uuid"
)

type NewUserPreferenceParams struct {
	UserID        uuid.UUID
	Theme         string
	Language      string
	Timezone      string
	Notifications map[string]interface{}
	Privacy       map[string]interface{}
	Display       map[string]interface{}
	Other         map[string]interface{}
}

func NewUserPreference(p NewUserPreferenceParams) (*UserPreference, error) {
	if err := validateUserPreferenceParams(p); err != nil {
		return nil, err
	}

	theme := p.Theme
	if theme == "" {
		theme = "light"
	}

	language := p.Language
	if language == "" {
		language = "en"
	}

	timezone := p.Timezone
	if timezone == "" {
		timezone = "UTC"
	}

	// id 必须等于 UserID（一对一关系，id 直接引用 users.id）
	now := time.Now().UTC()
	return NewUserPreferenceFromState(UserPreferenceState{
		ID:            p.UserID, // id 直接引用 users.id
		UserID:        p.UserID,
		Theme:         theme,
		Language:      language,
		Timezone:      timezone,
		Notifications: p.Notifications,
		Privacy:       p.Privacy,
		Display:       p.Display,
		Other:         p.Other,
		CreatedAt:     now,
		UpdatedAt:     now,
	}), nil
}

func NewUserPreferenceFromState(st UserPreferenceState) *UserPreference {
	return &UserPreference{state: st}
}

func validateUserPreferenceParams(p NewUserPreferenceParams) error {
	if p.UserID == uuid.Nil {
		return ErrUserIDRequired
	}
	return nil
}
