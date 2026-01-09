package mapper

import (
	"encoding/json"
	"nfxid/modules/directory/domain/user_preferences"
	"nfxid/modules/directory/infrastructure/rdb/models"
	"nfxid/pkgs/utils/timex"

	"gorm.io/datatypes"
)

// UserPreferenceDomainToModel 将 Domain UserPreference 转换为 Model UserPreference
func UserPreferenceDomainToModel(up *user_preferences.UserPreference) *models.UserPreference {
	if up == nil {
		return nil
	}

	var theme *string
	if up.Theme() != "" {
		t := up.Theme()
		theme = &t
	}

	var language *string
	if up.Language() != "" {
		l := up.Language()
		language = &l
	}

	var timezone *string
	if up.Timezone() != "" {
		tz := up.Timezone()
		timezone = &tz
	}

	var notifications *datatypes.JSON
	if up.Notifications() != nil && len(up.Notifications()) > 0 {
		notifBytes, _ := json.Marshal(up.Notifications())
		jsonData := datatypes.JSON(notifBytes)
		notifications = &jsonData
	}

	var privacy *datatypes.JSON
	if up.Privacy() != nil && len(up.Privacy()) > 0 {
		privBytes, _ := json.Marshal(up.Privacy())
		jsonData := datatypes.JSON(privBytes)
		privacy = &jsonData
	}

	var display *datatypes.JSON
	if up.Display() != nil && len(up.Display()) > 0 {
		dispBytes, _ := json.Marshal(up.Display())
		jsonData := datatypes.JSON(dispBytes)
		display = &jsonData
	}

	var other *datatypes.JSON
	if up.Other() != nil && len(up.Other()) > 0 {
		otherBytes, _ := json.Marshal(up.Other())
		jsonData := datatypes.JSON(otherBytes)
		other = &jsonData
	}

	return &models.UserPreference{
		ID:            up.ID(),
		UserID:        up.UserID(),
		Theme:         theme,
		Language:      language,
		Timezone:      timezone,
		Notifications: notifications,
		Privacy:       privacy,
		Display:       display,
		Other:         other,
		CreatedAt:     up.CreatedAt(),
		UpdatedAt:     up.UpdatedAt(),
		DeletedAt:     timex.TimeToGormDeletedAt(up.DeletedAt()),
	}
}

// UserPreferenceModelToDomain 将 Model UserPreference 转换为 Domain UserPreference
func UserPreferenceModelToDomain(m *models.UserPreference) *user_preferences.UserPreference {
	if m == nil {
		return nil
	}

	theme := ""
	if m.Theme != nil {
		theme = *m.Theme
	}

	language := ""
	if m.Language != nil {
		language = *m.Language
	}

	timezone := ""
	if m.Timezone != nil {
		timezone = *m.Timezone
	}

	var notifications map[string]interface{}
	if m.Notifications != nil {
		json.Unmarshal(*m.Notifications, &notifications)
	}

	var privacy map[string]interface{}
	if m.Privacy != nil {
		json.Unmarshal(*m.Privacy, &privacy)
	}

	var display map[string]interface{}
	if m.Display != nil {
		json.Unmarshal(*m.Display, &display)
	}

	var other map[string]interface{}
	if m.Other != nil {
		json.Unmarshal(*m.Other, &other)
	}

	state := user_preferences.UserPreferenceState{
		ID:            m.ID,
		UserID:        m.UserID,
		Theme:         theme,
		Language:      language,
		Timezone:      timezone,
		Notifications: notifications,
		Privacy:       privacy,
		Display:       display,
		Other:         other,
		CreatedAt:     m.CreatedAt,
		UpdatedAt:     m.UpdatedAt,
		DeletedAt:     timex.GormDeletedAtToTime(m.DeletedAt),
	}

	return user_preferences.NewUserPreferenceFromState(state)
}

// UserPreferenceModelToUpdates 将 Model UserPreference 转换为更新字段映射
func UserPreferenceModelToUpdates(m *models.UserPreference) map[string]any {
	var notifications any
	if m.Notifications != nil {
		notifications = m.Notifications
	}

	var privacy any
	if m.Privacy != nil {
		privacy = m.Privacy
	}

	var display any
	if m.Display != nil {
		display = m.Display
	}

	var other any
	if m.Other != nil {
		other = m.Other
	}

	return map[string]any{
		models.UserPreferenceCols.UserID:        m.UserID,
		models.UserPreferenceCols.Theme:         m.Theme,
		models.UserPreferenceCols.Language:      m.Language,
		models.UserPreferenceCols.Timezone:      m.Timezone,
		models.UserPreferenceCols.Notifications: notifications,
		models.UserPreferenceCols.Privacy:       privacy,
		models.UserPreferenceCols.Display:       display,
		models.UserPreferenceCols.Other:         other,
		models.UserPreferenceCols.UpdatedAt:     m.UpdatedAt,
		models.UserPreferenceCols.DeletedAt:     m.DeletedAt,
	}
}
