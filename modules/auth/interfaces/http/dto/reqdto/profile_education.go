package reqdto

import (
	"time"

	educationAppCommands "nfxid/modules/auth/application/profile_education/commands"
	educationAppQueries "nfxid/modules/auth/application/profile_education/queries"
	educationDomain "nfxid/modules/auth/domain/profile_education"
	"nfxid/pkgs/query"
	"nfxid/pkgs/utils/ptr"

	"github.com/google/uuid"
)

type EducationCreateRequestDTO struct {
	ProfileID    uuid.UUID  `json:"profile_id" validate:"required"`
	School       string     `json:"school" validate:"required"`
	Degree       *string    `json:"degree,omitempty"`
	Major        *string    `json:"major,omitempty"`
	FieldOfStudy *string    `json:"field_of_study,omitempty"`
	StartDate    *time.Time `json:"start_date,omitempty"`
	EndDate      *time.Time `json:"end_date,omitempty"`
	IsCurrent    bool       `json:"is_current"`
	Description  *string    `json:"description,omitempty"`
	Grade        *string    `json:"grade,omitempty"`
	Activities   *string    `json:"activities,omitempty"`
	Achievements *string    `json:"achievements,omitempty"`
}

type EducationUpdateRequestDTO struct {
	ID           uuid.UUID  `params:"id" validate:"required,uuid"`
	School       *string    `json:"school,omitempty"`
	Degree       *string    `json:"degree,omitempty"`
	Major        *string    `json:"major,omitempty"`
	FieldOfStudy *string    `json:"field_of_study,omitempty"`
	StartDate    *time.Time `json:"start_date,omitempty"`
	EndDate      *time.Time `json:"end_date,omitempty"`
	IsCurrent    *bool      `json:"is_current,omitempty"`
	Description  *string    `json:"description,omitempty"`
	Grade        *string    `json:"grade,omitempty"`
	Activities   *string    `json:"activities,omitempty"`
	Achievements *string    `json:"achievements,omitempty"`
}

type EducationByIDRequestDTO struct {
	ID uuid.UUID `params:"id" validate:"required,uuid"`
}

type EducationByProfileIDRequestDTO struct {
	ProfileID uuid.UUID `params:"profile_id" validate:"required,uuid"`
}

type EducationQueryParamsDTO struct {
	Offset    *int       `query:"offset"`
	Limit     *int       `query:"limit"`
	ProfileID *uuid.UUID `query:"profile_id"`
	Search    *string    `query:"search"`
	IsCurrent *bool      `query:"is_current"`
	Sort      []string   `query:"sort"`
}

func (r *EducationCreateRequestDTO) ToCreateCmd() educationAppCommands.CreateEducationCmd {
	return educationAppCommands.CreateEducationCmd{
		ProfileID: r.ProfileID,
		Editable: educationDomain.EducationEditable{
			School:       r.School,
			Degree:       r.Degree,
			Major:        r.Major,
			FieldOfStudy: r.FieldOfStudy,
			StartDate:    r.StartDate,
			EndDate:      r.EndDate,
			IsCurrent:    r.IsCurrent,
			Description:  r.Description,
			Grade:        r.Grade,
			Activities:   r.Activities,
			Achievements: r.Achievements,
		},
	}
}

func (r *EducationUpdateRequestDTO) ToUpdateCmd() educationAppCommands.UpdateEducationCmd {
	editable := educationDomain.EducationEditable{}
	if r.School != nil {
		editable.School = *r.School
	}
	if r.Degree != nil {
		editable.Degree = r.Degree
	}
	if r.Major != nil {
		editable.Major = r.Major
	}
	if r.FieldOfStudy != nil {
		editable.FieldOfStudy = r.FieldOfStudy
	}
	if r.StartDate != nil {
		editable.StartDate = r.StartDate
	}
	if r.EndDate != nil {
		editable.EndDate = r.EndDate
	}
	if r.IsCurrent != nil {
		editable.IsCurrent = *r.IsCurrent
	}
	if r.Description != nil {
		editable.Description = r.Description
	}
	if r.Grade != nil {
		editable.Grade = r.Grade
	}
	if r.Activities != nil {
		editable.Activities = r.Activities
	}
	if r.Achievements != nil {
		editable.Achievements = r.Achievements
	}

	return educationAppCommands.UpdateEducationCmd{
		EducationID: r.ID,
		Editable:    editable,
	}
}

func (r *EducationQueryParamsDTO) ToListQuery() educationAppQueries.EducationListQuery {
	var profileIDs []uuid.UUID
	if r.ProfileID != nil {
		profileIDs = []uuid.UUID{*r.ProfileID}
	}

	return educationAppQueries.EducationListQuery{
		DomainPagination: query.DomainPagination{
			Offset: ptr.Deref(r.Offset),
			Limit:  ptr.Deref(r.Limit),
		},
		DomainSorts: query.ParseSortParams(r.Sort, map[string]educationAppQueries.SortField{
			"created_time": educationAppQueries.SortByCreatedTime,
			"start_date":   educationAppQueries.SortByStartDate,
			"school":       educationAppQueries.SortBySchool,
		}),
		Search:     r.Search,
		ProfileIDs: profileIDs,
	}
}
