package mapper

import (
	"encoding/json"
	"nebulaid/modules/auth/domain/profile"
	"nebulaid/modules/auth/infrastructure/rdb/models"
	"nebulaid/pkgs/utils/timex"
	"strings"

	"github.com/google/uuid"
	"gorm.io/datatypes"
)

func ProfileDomainToModel(p *profile.Profile) *models.Profile {
	if p == nil {
		return nil
	}

	editable := p.Editable()

	// 序列化 JSON 字段
	var socialLinks *datatypes.JSON
	if editable.SocialLinks != nil {
		data, _ := json.Marshal(editable.SocialLinks)
		jsonData := datatypes.JSON(data)
		socialLinks = &jsonData
	}

	var preferences *datatypes.JSON
	if editable.Preferences != nil {
		data, _ := json.Marshal(editable.Preferences)
		jsonData := datatypes.JSON(data)
		preferences = &jsonData
	}

	var skills *datatypes.JSON
	if editable.Skills != nil {
		skillsMap := make(map[string]interface{})
		for k, v := range editable.Skills {
			skillsMap[k] = v
		}
		data, _ := json.Marshal(skillsMap)
		jsonData := datatypes.JSON(data)
		skills = &jsonData
	}

	var privacySettings *datatypes.JSON
	if editable.PrivacySettings != nil {
		data, _ := json.Marshal(editable.PrivacySettings)
		jsonData := datatypes.JSON(data)
		privacySettings = &jsonData
	}

	// 序列化 BackgroundIds 为 PostgreSQL UUID数组格式 "{uuid1,uuid2,...}"
	var backgroundIdsStr *string
	if len(editable.BackgroundIds) > 0 {
		ids := make([]string, len(editable.BackgroundIds))
		for i, id := range editable.BackgroundIds {
			ids[i] = id.String()
		}
		idsStr := "{" + strings.Join(ids, ",") + "}"
		backgroundIdsStr = &idsStr
	}

	return &models.Profile{
		ID:              p.ID(),
		UserID:          p.UserID(),
		FirstName:       editable.FirstName,
		LastName:        editable.LastName,
		Nickname:        editable.Nickname,
		DisplayName:     editable.DisplayName,
		AvatarID:        editable.AvatarID,
		BackgroundID:    editable.BackgroundID,
		BackgroundIds:   backgroundIdsStr,
		Bio:             editable.Bio,
		Phone:           editable.Phone,
		Birthday:        editable.Birthday,
		Age:             editable.Age,
		Gender:          editable.Gender,
		Location:        editable.Location,
		Website:         editable.Website,
		Github:          editable.Github,
		SocialLinks:     socialLinks,
		Preferences:     preferences,
		Skills:          skills,
		PrivacySettings: privacySettings,
		CreatedAt:       p.CreatedAt(),
		UpdatedAt:       p.UpdatedAt(),
		DeletedAt:       timex.TimeToGormDeletedAt(p.DeletedAt()),
	}
}

func ProfileModelToDomain(m *models.Profile) *profile.Profile {
	if m == nil {
		return nil
	}

	// 解析 JSON 字段
	var socialLinks map[string]interface{}
	if m.SocialLinks != nil {
		json.Unmarshal(*m.SocialLinks, &socialLinks)
	}

	var preferences map[string]interface{}
	if m.Preferences != nil {
		json.Unmarshal(*m.Preferences, &preferences)
	}

	var skills map[string]int
	if m.Skills != nil {
		var skillsMap map[string]interface{}
		json.Unmarshal(*m.Skills, &skillsMap)
		skills = make(map[string]int)
		for k, v := range skillsMap {
			if intVal, ok := v.(float64); ok {
				skills[k] = int(intVal)
			}
		}
	}

	var privacySettings map[string]interface{}
	if m.PrivacySettings != nil {
		json.Unmarshal(*m.PrivacySettings, &privacySettings)
	}

	// 解析 BackgroundIds (UUID数组)
	var backgroundIds []uuid.UUID
	if m.BackgroundIds != nil && *m.BackgroundIds != "" {
		idsStr := strings.Trim(*m.BackgroundIds, "{}")
		if idsStr != "" {
			parts := strings.Split(idsStr, ",")
			backgroundIds = make([]uuid.UUID, 0, len(parts))
			for _, part := range parts {
				part = strings.TrimSpace(part)
				if id, err := uuid.Parse(part); err == nil {
					backgroundIds = append(backgroundIds, id)
				}
			}
		}
	}

	editable := profile.ProfileEditable{
		FirstName:       m.FirstName,
		LastName:        m.LastName,
		Nickname:        m.Nickname,
		DisplayName:     m.DisplayName,
		AvatarID:        m.AvatarID,
		BackgroundID:    m.BackgroundID,
		BackgroundIds:   backgroundIds,
		Bio:             m.Bio,
		Phone:           m.Phone,
		Birthday:        m.Birthday,
		Age:             m.Age,
		Gender:          m.Gender,
		Location:        m.Location,
		Website:         m.Website,
		Github:          m.Github,
		SocialLinks:     socialLinks,
		Preferences:     preferences,
		Skills:          skills,
		PrivacySettings: privacySettings,
	}

	state := profile.ProfileState{
		ID:        m.ID,
		UserID:    m.UserID,
		Editable:  editable,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
		DeletedAt: timex.GormDeletedAtToTime(m.DeletedAt),
	}

	return profile.NewProfileFromState(state)
}

func ProfileModelsToUpdates(m *models.Profile) map[string]any {
	return map[string]any{
		models.ProfileCols.FirstName:       m.FirstName,
		models.ProfileCols.LastName:        m.LastName,
		models.ProfileCols.Nickname:        m.Nickname,
		models.ProfileCols.DisplayName:     m.DisplayName,
		models.ProfileCols.AvatarID:        m.AvatarID,
		models.ProfileCols.BackgroundID:    m.BackgroundID,
		models.ProfileCols.BackgroundIds:   m.BackgroundIds,
		models.ProfileCols.Bio:             m.Bio,
		models.ProfileCols.Phone:           m.Phone,
		models.ProfileCols.Birthday:        m.Birthday,
		models.ProfileCols.Age:             m.Age,
		models.ProfileCols.Gender:          m.Gender,
		models.ProfileCols.Location:        m.Location,
		models.ProfileCols.Website:         m.Website,
		models.ProfileCols.Github:          m.Github,
		models.ProfileCols.SocialLinks:     m.SocialLinks,
		models.ProfileCols.Preferences:     m.Preferences,
		models.ProfileCols.Skills:          m.Skills,
		models.ProfileCols.PrivacySettings: m.PrivacySettings,
		models.ProfileCols.DeletedAt:       m.DeletedAt,
	}
}
