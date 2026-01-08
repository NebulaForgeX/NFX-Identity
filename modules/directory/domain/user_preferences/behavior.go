package user_preferences

import (
	"time"
)

func (up *UserPreference) Update(theme, language, timezone *string, notifications, privacy, display, other map[string]interface{}) error {
	if up.DeletedAt() != nil {
		return ErrUserPreferenceNotFound
	}

	if theme != nil && *theme != "" {
		up.state.Theme = *theme
	}
	if language != nil && *language != "" {
		up.state.Language = *language
	}
	if timezone != nil && *timezone != "" {
		up.state.Timezone = *timezone
	}
	if notifications != nil {
		up.state.Notifications = notifications
	}
	if privacy != nil {
		up.state.Privacy = privacy
	}
	if display != nil {
		up.state.Display = display
	}
	if other != nil {
		up.state.Other = other
	}

	up.state.UpdatedAt = time.Now().UTC()
	return nil
}

func (up *UserPreference) UpdateTheme(theme string) error {
	if up.DeletedAt() != nil {
		return ErrUserPreferenceNotFound
	}
	up.state.Theme = theme
	up.state.UpdatedAt = time.Now().UTC()
	return nil
}

func (up *UserPreference) UpdateLanguage(language string) error {
	if up.DeletedAt() != nil {
		return ErrUserPreferenceNotFound
	}
	up.state.Language = language
	up.state.UpdatedAt = time.Now().UTC()
	return nil
}

func (up *UserPreference) UpdateTimezone(timezone string) error {
	if up.DeletedAt() != nil {
		return ErrUserPreferenceNotFound
	}
	up.state.Timezone = timezone
	up.state.UpdatedAt = time.Now().UTC()
	return nil
}

func (up *UserPreference) Delete() error {
	if up.DeletedAt() != nil {
		return nil // idempotent
	}

	now := time.Now().UTC()
	up.state.DeletedAt = &now
	up.state.UpdatedAt = now
	return nil
}
