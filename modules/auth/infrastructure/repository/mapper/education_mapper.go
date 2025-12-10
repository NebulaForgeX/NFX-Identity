package mapper

import (
	"nebulaid/modules/auth/domain/education"
	"nebulaid/modules/auth/infrastructure/rdb/models"
	"nebulaid/pkgs/utils/timex"
)

func EducationDomainToModel(e *education.Education) *models.Education {
	if e == nil {
		return nil
	}

	editable := e.Editable()
	return &models.Education{
		ID:           e.ID(),
		ProfileID:    e.ProfileID(),
		School:       editable.School,
		Degree:       editable.Degree,
		Major:        editable.Major,
		FieldOfStudy: editable.FieldOfStudy,
		StartDate:    editable.StartDate,
		EndDate:      editable.EndDate,
		IsCurrent:    editable.IsCurrent,
		Description:  editable.Description,
		Grade:        editable.Grade,
		Activities:   editable.Activities,
		Achievements: editable.Achievements,
		CreatedAt:    e.CreatedAt(),
		UpdatedAt:    e.UpdatedAt(),
		DeletedAt:    timex.TimeToGormDeletedAt(e.DeletedAt()),
	}
}

func EducationModelToDomain(m *models.Education) *education.Education {
	if m == nil {
		return nil
	}

	editable := education.EducationEditable{
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
	}

	state := education.EducationState{
		ID:        m.ID,
		ProfileID: m.ProfileID,
		Editable:  editable,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
		DeletedAt: timex.GormDeletedAtToTime(m.DeletedAt),
	}

	return education.NewEducationFromState(state)
}

func EducationModelsToUpdates(m *models.Education) map[string]any {
	return map[string]any{
		models.EducationCols.ProfileID:    m.ProfileID,
		models.EducationCols.School:       m.School,
		models.EducationCols.Degree:       m.Degree,
		models.EducationCols.Major:        m.Major,
		models.EducationCols.FieldOfStudy: m.FieldOfStudy,
		models.EducationCols.StartDate:    m.StartDate,
		models.EducationCols.EndDate:      m.EndDate,
		models.EducationCols.IsCurrent:    m.IsCurrent,
		models.EducationCols.Description:  m.Description,
		models.EducationCols.Grade:        m.Grade,
		models.EducationCols.Activities:   m.Activities,
		models.EducationCols.Achievements: m.Achievements,
		models.EducationCols.DeletedAt:    m.DeletedAt,
	}
}
