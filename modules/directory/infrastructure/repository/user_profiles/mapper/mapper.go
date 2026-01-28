package mapper

import (
	"encoding/json"
	"nfxid/modules/directory/domain/user_profiles"
	"nfxid/modules/directory/infrastructure/rdb/models"
	"nfxid/pkgs/utils/timex"

	"gorm.io/datatypes"
)

// UserProfileDomainToModel 将 Domain UserProfile 转换为 Model UserProfile
func UserProfileDomainToModel(up *user_profiles.UserProfile) *models.UserProfile {
	if up == nil {
		return nil
	}

	var socialLinks *datatypes.JSON
	if up.SocialLinks() != nil && len(up.SocialLinks()) > 0 {
		linksBytes, _ := json.Marshal(up.SocialLinks())
		jsonData := datatypes.JSON(linksBytes)
		socialLinks = &jsonData
	}

	var skills *datatypes.JSON
	if up.Skills() != nil && len(up.Skills()) > 0 {
		skillsBytes, _ := json.Marshal(up.Skills())
		jsonData := datatypes.JSON(skillsBytes)
		skills = &jsonData
	}

	return &models.UserProfile{
		ID:          up.ID(), // id 直接引用 users.id
		Role:        up.Role(),
		FirstName:   up.FirstName(),
		LastName:    up.LastName(),
		Nickname:    up.Nickname(),
		DisplayName: up.DisplayName(),
		Bio:         up.Bio(),
		Birthday:    up.Birthday(),
		Age:         up.Age(),
		Gender:      up.Gender(),
		Location:    up.Location(),
		Website:     up.Website(),
		Github:      up.Github(),
		SocialLinks: socialLinks,
		Skills:      skills,
		CreatedAt:   up.CreatedAt(),
		UpdatedAt:   up.UpdatedAt(),
		DeletedAt:   timex.TimeToGormDeletedAt(up.DeletedAt()),
	}
}

// UserProfileModelToDomain 将 Model UserProfile 转换为 Domain UserProfile
func UserProfileModelToDomain(m *models.UserProfile) *user_profiles.UserProfile {
	if m == nil {
		return nil
	}

	var socialLinks map[string]interface{}
	if m.SocialLinks != nil {
		json.Unmarshal(*m.SocialLinks, &socialLinks)
	}

	var skills map[string]interface{}
	if m.Skills != nil {
		json.Unmarshal(*m.Skills, &skills)
	}

	state := user_profiles.UserProfileState{
		ID:          m.ID, // id 直接引用 users.id
		UserID:      m.ID, // UserID 从 ID 获取（一对一关系）
		Role:        m.Role,
		FirstName:   m.FirstName,
		LastName:    m.LastName,
		Nickname:    m.Nickname,
		DisplayName: m.DisplayName,
		Bio:         m.Bio,
		Birthday:    m.Birthday,
		Age:         m.Age,
		Gender:      m.Gender,
		Location:    m.Location,
		Website:     m.Website,
		Github:      m.Github,
		SocialLinks: socialLinks,
		Skills:      skills,
		CreatedAt:   m.CreatedAt,
		UpdatedAt:   m.UpdatedAt,
		DeletedAt:   timex.GormDeletedAtToTime(m.DeletedAt),
	}

	return user_profiles.NewUserProfileFromState(state)
}

// UserProfileModelToUpdates 将 Model UserProfile 转换为更新字段映射
func UserProfileModelToUpdates(m *models.UserProfile) map[string]any {
	var socialLinks any
	if m.SocialLinks != nil {
		socialLinks = m.SocialLinks
	}

	var skills any
	if m.Skills != nil {
		skills = m.Skills
	}

	return map[string]any{
		// 注意：UserID 不再存在，id 直接引用 users.id
		models.UserProfileCols.Role:        m.Role,
		models.UserProfileCols.FirstName:   m.FirstName,
		models.UserProfileCols.LastName:    m.LastName,
		models.UserProfileCols.Nickname:    m.Nickname,
		models.UserProfileCols.DisplayName: m.DisplayName,
		models.UserProfileCols.Bio:         m.Bio,
		models.UserProfileCols.Birthday:    m.Birthday,
		models.UserProfileCols.Age:         m.Age,
		models.UserProfileCols.Gender:      m.Gender,
		models.UserProfileCols.Location:    m.Location,
		models.UserProfileCols.Website:     m.Website,
		models.UserProfileCols.Github:      m.Github,
		models.UserProfileCols.SocialLinks: socialLinks,
		models.UserProfileCols.Skills:      skills,
		models.UserProfileCols.UpdatedAt:   m.UpdatedAt,
		models.UserProfileCols.DeletedAt:   m.DeletedAt,
	}
}
