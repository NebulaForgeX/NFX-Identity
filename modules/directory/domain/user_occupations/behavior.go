package user_occupations

import (
	"time"
)

func (uo *UserOccupation) Update(company, position string, department, industry, location, employmentType *string, startDate, endDate *time.Time, isCurrent bool, description, responsibilities, achievements *string, skillsUsed []string) error {
	if uo.DeletedAt() != nil {
		return ErrUserOccupationNotFound
	}
	if company == "" {
		return ErrCompanyRequired
	}
	if position == "" {
		return ErrPositionRequired
	}

	uo.state.Company = company
	uo.state.Position = position
	if department != nil {
		uo.state.Department = department
	}
	if industry != nil {
		uo.state.Industry = industry
	}
	if location != nil {
		uo.state.Location = location
	}
	if employmentType != nil {
		uo.state.EmploymentType = employmentType
	}
	if startDate != nil {
		uo.state.StartDate = startDate
	}
	if endDate != nil {
		uo.state.EndDate = endDate
	}
	uo.state.IsCurrent = isCurrent
	if description != nil {
		uo.state.Description = description
	}
	if responsibilities != nil {
		uo.state.Responsibilities = responsibilities
	}
	if achievements != nil {
		uo.state.Achievements = achievements
	}
	if skillsUsed != nil {
		uo.state.SkillsUsed = skillsUsed
	}

	uo.state.UpdatedAt = time.Now().UTC()
	return nil
}

func (uo *UserOccupation) Delete() error {
	if uo.DeletedAt() != nil {
		return nil // idempotent
	}

	now := time.Now().UTC()
	uo.state.DeletedAt = &now
	uo.state.UpdatedAt = now
	return nil
}
