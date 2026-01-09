package mapper

import (
	"strings"
	"nfxid/modules/directory/domain/user_occupations"
	"nfxid/modules/directory/infrastructure/rdb/models"
	"nfxid/pkgs/utils/timex"
)

// UserOccupationDomainToModel 将 Domain UserOccupation 转换为 Model UserOccupation
func UserOccupationDomainToModel(uo *user_occupations.UserOccupation) *models.UserOccupation {
	if uo == nil {
		return nil
	}

	// 序列化 SkillsUsed 为 PostgreSQL text[] 格式 "{skill1,skill2,...}"
	var skillsUsedStr *string
	if len(uo.SkillsUsed()) > 0 {
		// PostgreSQL text[] 格式需要转义特殊字符，但简单情况下可以直接用逗号分隔
		skillsStr := "{" + strings.Join(uo.SkillsUsed(), ",") + "}"
		skillsUsedStr = &skillsStr
	}

	return &models.UserOccupation{
		ID:               uo.ID(),
		UserID:           uo.UserID(),
		Company:          uo.Company(),
		Position:         uo.Position(),
		Department:       uo.Department(),
		Industry:          uo.Industry(),
		Location:          uo.Location(),
		EmploymentType:    uo.EmploymentType(),
		StartDate:         uo.StartDate(),
		EndDate:           uo.EndDate(),
		IsCurrent:         uo.IsCurrent(),
		Description:       uo.Description(),
		Responsibilities:  uo.Responsibilities(),
		Achievements:      uo.Achievements(),
		SkillsUsed:         skillsUsedStr,
		CreatedAt:          uo.CreatedAt(),
		UpdatedAt:          uo.UpdatedAt(),
		DeletedAt:          timex.TimeToGormDeletedAt(uo.DeletedAt()),
	}
}

// UserOccupationModelToDomain 将 Model UserOccupation 转换为 Domain UserOccupation
func UserOccupationModelToDomain(m *models.UserOccupation) *user_occupations.UserOccupation {
	if m == nil {
		return nil
	}

	// 解析 SkillsUsed 从 PostgreSQL text[] 格式 "{skill1,skill2,...}"
	var skillsUsed []string
	if m.SkillsUsed != nil && *m.SkillsUsed != "" {
		// 移除大括号并分割
		skillsStr := strings.Trim(*m.SkillsUsed, "{}")
		if skillsStr != "" {
			skillsUsed = strings.Split(skillsStr, ",")
			// 去除每个技能的空格
			for i, skill := range skillsUsed {
				skillsUsed[i] = strings.TrimSpace(skill)
			}
		}
	}

	state := user_occupations.UserOccupationState{
		ID:               m.ID,
		UserID:           m.UserID,
		Company:          m.Company,
		Position:         m.Position,
		Department:       m.Department,
		Industry:          m.Industry,
		Location:          m.Location,
		EmploymentType:    m.EmploymentType,
		StartDate:         m.StartDate,
		EndDate:           m.EndDate,
		IsCurrent:         m.IsCurrent,
		Description:       m.Description,
		Responsibilities:  m.Responsibilities,
		Achievements:      m.Achievements,
		SkillsUsed:         skillsUsed,
		CreatedAt:          m.CreatedAt,
		UpdatedAt:          m.UpdatedAt,
		DeletedAt:          timex.GormDeletedAtToTime(m.DeletedAt),
	}

	return user_occupations.NewUserOccupationFromState(state)
}

// UserOccupationModelToUpdates 将 Model UserOccupation 转换为更新字段映射
func UserOccupationModelToUpdates(m *models.UserOccupation) map[string]any {
	return map[string]any{
		models.UserOccupationCols.UserID:           m.UserID,
		models.UserOccupationCols.Company:          m.Company,
		models.UserOccupationCols.Position:         m.Position,
		models.UserOccupationCols.Department:       m.Department,
		models.UserOccupationCols.Industry:          m.Industry,
		models.UserOccupationCols.Location:          m.Location,
		models.UserOccupationCols.EmploymentType:    m.EmploymentType,
		models.UserOccupationCols.StartDate:         m.StartDate,
		models.UserOccupationCols.EndDate:           m.EndDate,
		models.UserOccupationCols.IsCurrent:         m.IsCurrent,
		models.UserOccupationCols.Description:       m.Description,
		models.UserOccupationCols.Responsibilities:  m.Responsibilities,
		models.UserOccupationCols.Achievements:      m.Achievements,
		models.UserOccupationCols.SkillsUsed:         m.SkillsUsed,
		models.UserOccupationCols.UpdatedAt:          m.UpdatedAt,
		models.UserOccupationCols.DeletedAt:          m.DeletedAt,
	}
}
