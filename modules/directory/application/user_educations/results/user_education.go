package results

import (
	"time"

	"nfxid/modules/directory/domain/user_educations"

	"github.com/google/uuid"
)

type UserEducationRO struct {
	ID          uuid.UUID
	UserID      uuid.UUID
	School      string
	Degree      *string
	Major       *string
	FieldOfStudy *string
	StartDate   *time.Time
	EndDate     *time.Time
	IsCurrent   bool
	Description *string
	Grade       *string
	Activities  *string
	Achievements *string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
}

// UserEducationMapper 将 Domain UserEducation 转换为 Application UserEducationRO
func UserEducationMapper(ue *user_educations.UserEducation) UserEducationRO {
	if ue == nil {
		return UserEducationRO{}
	}

	return UserEducationRO{
		ID:          ue.ID(),
		UserID:      ue.UserID(),
		School:      ue.School(),
		Degree:      ue.Degree(),
		Major:       ue.Major(),
		FieldOfStudy: ue.FieldOfStudy(),
		StartDate:   ue.StartDate(),
		EndDate:     ue.EndDate(),
		IsCurrent:   ue.IsCurrent(),
		Description: ue.Description(),
		Grade:       ue.Grade(),
		Activities:  ue.Activities(),
		Achievements: ue.Achievements(),
		CreatedAt:   ue.CreatedAt(),
		UpdatedAt:   ue.UpdatedAt(),
		DeletedAt:   ue.DeletedAt(),
	}
}
