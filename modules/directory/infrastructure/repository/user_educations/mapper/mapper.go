package mapper

import (
	"nfxid/modules/directory/domain/user_educations"
	"nfxid/modules/directory/infrastructure/rdb/models"
	"nfxid/pkgs/utils/timex"
)

// UserEducationDomainToModel 将 Domain UserEducation 转换为 Model UserEducation
func UserEducationDomainToModel(ue *user_educations.UserEducation) *models.UserEducation {
	if ue == nil {
		return nil
	}

	return &models.UserEducation{
		ID:           ue.ID(),
		UserID:       ue.UserID(),
		School:       ue.School(),
		Degree:       ue.Degree(),
		Major:        ue.Major(),
		FieldOfStudy: ue.FieldOfStudy(),
		StartDate:    ue.StartDate(),
		EndDate:      ue.EndDate(),
		IsCurrent:    ue.IsCurrent(),
		Description:  ue.Description(),
		Grade:        ue.Grade(),
		Activities:   ue.Activities(),
		Achievements: ue.Achievements(),
		CreatedAt:    ue.CreatedAt(),
		UpdatedAt:    ue.UpdatedAt(),
		DeletedAt:   timex.TimeToGormDeletedAt(ue.DeletedAt()),
	}
}

// UserEducationModelToDomain 将 Model UserEducation 转换为 Domain UserEducation
func UserEducationModelToDomain(m *models.UserEducation) *user_educations.UserEducation {
	if m == nil {
		return nil
	}

	state := user_educations.UserEducationState{
		ID:           m.ID,
		UserID:       m.UserID,
		School:       m.School,
		Degree:       m.Degree,
		Major:        m.Major,
		FieldOfStudy: m.FieldOfStudy,
		StartDate:    m.StartDate,
		EndDate:      m.EndDate,
		IsCurrent:    m.IsCurrent,
		Description:  m.Description,
		Grade:        m.Grade,
		Activities:   m.Activities,
		Achievements: m.Achievements,
		CreatedAt:    m.CreatedAt,
		UpdatedAt:    m.UpdatedAt,
		DeletedAt:    timex.GormDeletedAtToTime(m.DeletedAt),
	}

	return user_educations.NewUserEducationFromState(state)
}

// UserEducationModelToUpdates 将 Model UserEducation 转换为更新字段映射
func UserEducationModelToUpdates(m *models.UserEducation) map[string]any {
	return map[string]any{
		models.UserEducationCols.UserID:       m.UserID,
		models.UserEducationCols.School:       m.School,
		models.UserEducationCols.Degree:       m.Degree,
		models.UserEducationCols.Major:        m.Major,
		models.UserEducationCols.FieldOfStudy: m.FieldOfStudy,
		models.UserEducationCols.StartDate:    m.StartDate,
		models.UserEducationCols.EndDate:      m.EndDate,
		models.UserEducationCols.IsCurrent:    m.IsCurrent,
		models.UserEducationCols.Description:  m.Description,
		models.UserEducationCols.Grade:        m.Grade,
		models.UserEducationCols.Activities:   m.Activities,
		models.UserEducationCols.Achievements: m.Achievements,
		models.UserEducationCols.UpdatedAt:    m.UpdatedAt,
		models.UserEducationCols.DeletedAt:   m.DeletedAt,
	}
}
