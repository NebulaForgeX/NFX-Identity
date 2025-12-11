package reqdto

import (
	"time"

	occupationAppCommands "nfxid/modules/auth/application/occupation/commands"
	occupationAppQueries "nfxid/modules/auth/application/occupation/queries"
	occupationDomain "nfxid/modules/auth/domain/occupation"
	"nfxid/pkgs/query"
	"nfxid/pkgs/utils/ptr"

	"github.com/google/uuid"
)

type OccupationCreateRequestDTO struct {
	ProfileID        uuid.UUID  `json:"profile_id" validate:"required"`
	Company          string     `json:"company" validate:"required"`
	Position         string     `json:"position" validate:"required"`
	Department       *string    `json:"department,omitempty"`
	Industry         *string    `json:"industry,omitempty"`
	Location         *string    `json:"location,omitempty"`
	EmploymentType   *string    `json:"employment_type,omitempty"`
	StartDate        *time.Time `json:"start_date,omitempty"`
	EndDate          *time.Time `json:"end_date,omitempty"`
	IsCurrent        bool       `json:"is_current"`
	Description      *string    `json:"description,omitempty"`
	Responsibilities *string    `json:"responsibilities,omitempty"`
	Achievements     *string    `json:"achievements,omitempty"`
	SkillsUsed       []string   `json:"skills_used,omitempty"`
}

type OccupationUpdateRequestDTO struct {
	ID               uuid.UUID  `params:"id" validate:"required,uuid"`
	Company          *string    `json:"company,omitempty"`
	Position         *string    `json:"position,omitempty"`
	Department       *string    `json:"department,omitempty"`
	Industry         *string    `json:"industry,omitempty"`
	Location         *string    `json:"location,omitempty"`
	EmploymentType   *string    `json:"employment_type,omitempty"`
	StartDate        *time.Time `json:"start_date,omitempty"`
	EndDate          *time.Time `json:"end_date,omitempty"`
	IsCurrent        *bool      `json:"is_current,omitempty"`
	Description      *string    `json:"description,omitempty"`
	Responsibilities *string    `json:"responsibilities,omitempty"`
	Achievements     *string    `json:"achievements,omitempty"`
	SkillsUsed       []string   `json:"skills_used,omitempty"`
}

type OccupationByIDRequestDTO struct {
	ID uuid.UUID `params:"id" validate:"required,uuid"`
}

type OccupationByProfileIDRequestDTO struct {
	ProfileID uuid.UUID `params:"profile_id" validate:"required,uuid"`
}

type OccupationQueryParamsDTO struct {
	Offset    *int       `query:"offset"`
	Limit     *int       `query:"limit"`
	ProfileID *uuid.UUID `query:"profile_id"`
	Search    *string    `query:"search"`
	IsCurrent *bool      `query:"is_current"`
	Industry  *string    `query:"industry"`
	Sort      []string   `query:"sort"`
}

func (r *OccupationCreateRequestDTO) ToCreateCmd() occupationAppCommands.CreateOccupationCmd {
	return occupationAppCommands.CreateOccupationCmd{
		ProfileID: r.ProfileID,
		Editable: occupationDomain.OccupationEditable{
			Company:          r.Company,
			Position:         r.Position,
			Department:       r.Department,
			Industry:         r.Industry,
			Location:         r.Location,
			EmploymentType:   r.EmploymentType,
			StartDate:        r.StartDate,
			EndDate:          r.EndDate,
			IsCurrent:        r.IsCurrent,
			Description:      r.Description,
			Responsibilities: r.Responsibilities,
			Achievements:     r.Achievements,
			SkillsUsed:       r.SkillsUsed,
		},
	}
}

func (r *OccupationUpdateRequestDTO) ToUpdateCmd() occupationAppCommands.UpdateOccupationCmd {
	editable := occupationDomain.OccupationEditable{}
	if r.Company != nil {
		editable.Company = *r.Company
	}
	if r.Position != nil {
		editable.Position = *r.Position
	}
	if r.Department != nil {
		editable.Department = r.Department
	}
	if r.Industry != nil {
		editable.Industry = r.Industry
	}
	if r.Location != nil {
		editable.Location = r.Location
	}
	if r.EmploymentType != nil {
		editable.EmploymentType = r.EmploymentType
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
	if r.Responsibilities != nil {
		editable.Responsibilities = r.Responsibilities
	}
	if r.Achievements != nil {
		editable.Achievements = r.Achievements
	}
	if r.SkillsUsed != nil {
		editable.SkillsUsed = r.SkillsUsed
	}

	return occupationAppCommands.UpdateOccupationCmd{
		OccupationID: r.ID,
		Editable:     editable,
	}
}

func (r *OccupationQueryParamsDTO) ToListQuery() occupationAppQueries.OccupationListQuery {
	var profileIDs []uuid.UUID
	if r.ProfileID != nil {
		profileIDs = []uuid.UUID{*r.ProfileID}
	}

	return occupationAppQueries.OccupationListQuery{
		DomainPagination: query.DomainPagination{
			Offset: ptr.Deref(r.Offset),
			Limit:  ptr.Deref(r.Limit),
		},
		DomainSorts: query.ParseSortParams(r.Sort, map[string]occupationAppQueries.SortField{
			"created_time": occupationAppQueries.SortByCreatedTime,
			"start_date":   occupationAppQueries.SortByStartDate,
			"company":      occupationAppQueries.SortByCompany,
		}),
		Search:     r.Search,
		ProfileIDs: profileIDs,
	}
}
