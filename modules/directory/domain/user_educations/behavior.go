package user_educations

import (
	"time"
)

func (ue *UserEducation) Update(school string, degree, major, fieldOfStudy *string, startDate, endDate *time.Time, isCurrent bool, description, grade, activities, achievements *string) error {
	if ue.DeletedAt() != nil {
		return ErrUserEducationNotFound
	}
	if school == "" {
		return ErrSchoolRequired
	}

	ue.state.School = school
	if degree != nil {
		ue.state.Degree = degree
	}
	if major != nil {
		ue.state.Major = major
	}
	if fieldOfStudy != nil {
		ue.state.FieldOfStudy = fieldOfStudy
	}
	if startDate != nil {
		ue.state.StartDate = startDate
	}
	if endDate != nil {
		ue.state.EndDate = endDate
	}
	ue.state.IsCurrent = isCurrent
	if description != nil {
		ue.state.Description = description
	}
	if grade != nil {
		ue.state.Grade = grade
	}
	if activities != nil {
		ue.state.Activities = activities
	}
	if achievements != nil {
		ue.state.Achievements = achievements
	}

	ue.state.UpdatedAt = time.Now().UTC()
	return nil
}

func (ue *UserEducation) Delete() error {
	if ue.DeletedAt() != nil {
		return nil // idempotent
	}

	now := time.Now().UTC()
	ue.state.DeletedAt = &now
	ue.state.UpdatedAt = now
	return nil
}
