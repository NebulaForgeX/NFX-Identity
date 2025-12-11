package mapper

import (
	"nfxid/modules/auth/domain/occupation"
	"nfxid/modules/auth/infrastructure/rdb/models"
	"nfxid/pkgs/utils/timex"
	"strings"
)

func OccupationDomainToModel(o *occupation.Occupation) *models.Occupation {
	if o == nil {
		return nil
	}

	editable := o.Editable()

	// 格式化 SkillsUsed 为 PostgreSQL text[] 格式
	var skillsUsedStr *string
	if len(editable.SkillsUsed) > 0 {
		formatted := "{" + strings.Join(editable.SkillsUsed, ",") + "}"
		skillsUsedStr = &formatted
	}

	return &models.Occupation{
		ID:               o.ID(),
		ProfileID:        o.ProfileID(),
		Company:          editable.Company,
		Position:         editable.Position,
		Department:       editable.Department,
		Industry:         editable.Industry,
		Location:         editable.Location,
		EmploymentType:   editable.EmploymentType,
		StartDate:        editable.StartDate,
		EndDate:          editable.EndDate,
		IsCurrent:        editable.IsCurrent,
		Description:      editable.Description,
		Responsibilities: editable.Responsibilities,
		Achievements:     editable.Achievements,
		SkillsUsed:       skillsUsedStr,
		CreatedAt:        o.CreatedAt(),
		UpdatedAt:        o.UpdatedAt(),
		DeletedAt:        timex.TimeToGormDeletedAt(o.DeletedAt()),
	}
}

func OccupationModelToDomain(m *models.Occupation) *occupation.Occupation {
	if m == nil {
		return nil
	}

	// 解析 SkillsUsed (PostgreSQL text[] 格式)
	var skillsUsed []string
	if m.SkillsUsed != nil && *m.SkillsUsed != "" {
		trimmed := strings.Trim(*m.SkillsUsed, "{}")
		if trimmed != "" {
			skillsUsed = strings.Split(trimmed, ",")
		}
	}

	editable := occupation.OccupationEditable{
		Company:          m.Company,
		Position:         m.Position,
		Department:       m.Department,
		Industry:         m.Industry,
		Location:         m.Location,
		EmploymentType:   m.EmploymentType,
		StartDate:        m.StartDate,
		EndDate:          m.EndDate,
		IsCurrent:        m.IsCurrent,
		Description:      m.Description,
		Responsibilities: m.Responsibilities,
		Achievements:     m.Achievements,
		SkillsUsed:       skillsUsed,
	}

	state := occupation.OccupationState{
		ID:        m.ID,
		ProfileID: m.ProfileID,
		Editable:  editable,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
		DeletedAt: timex.GormDeletedAtToTime(m.DeletedAt),
	}

	return occupation.NewOccupationFromState(state)
}

func OccupationModelsToUpdates(m *models.Occupation) map[string]any {
	return map[string]any{
		models.OccupationCols.ProfileID:        m.ProfileID,
		models.OccupationCols.Company:          m.Company,
		models.OccupationCols.Position:         m.Position,
		models.OccupationCols.Department:       m.Department,
		models.OccupationCols.Industry:         m.Industry,
		models.OccupationCols.Location:         m.Location,
		models.OccupationCols.EmploymentType:   m.EmploymentType,
		models.OccupationCols.StartDate:        m.StartDate,
		models.OccupationCols.EndDate:          m.EndDate,
		models.OccupationCols.IsCurrent:        m.IsCurrent,
		models.OccupationCols.Description:      m.Description,
		models.OccupationCols.Responsibilities: m.Responsibilities,
		models.OccupationCols.Achievements:     m.Achievements,
		models.OccupationCols.SkillsUsed:       m.SkillsUsed,
		models.OccupationCols.DeletedAt:        m.DeletedAt,
	}
}
