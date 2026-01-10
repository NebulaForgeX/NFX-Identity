package user_educations

import (
	"context"
	"time"
	userEducationCommands "nfxid/modules/directory/application/user_educations/commands"
	userEducationDomain "nfxid/modules/directory/domain/user_educations"

	"github.com/google/uuid"
)

// CreateUserEducation 创建用户教育经历
func (s *Service) CreateUserEducation(ctx context.Context, cmd userEducationCommands.CreateUserEducationCmd) (uuid.UUID, error) {
	var startDate, endDate *time.Time
	if cmd.StartDate != nil && *cmd.StartDate != "" {
		parsed, err := time.Parse(time.RFC3339, *cmd.StartDate)
		if err != nil {
			return uuid.Nil, err
		}
		startDate = &parsed
	}
	if cmd.EndDate != nil && *cmd.EndDate != "" {
		parsed, err := time.Parse(time.RFC3339, *cmd.EndDate)
		if err != nil {
			return uuid.Nil, err
		}
		endDate = &parsed
	}

	// Create domain entity
	userEducation, err := userEducationDomain.NewUserEducation(userEducationDomain.NewUserEducationParams{
		UserID:      cmd.UserID,
		School:      cmd.School,
		Degree:      cmd.Degree,
		Major:       cmd.Major,
		FieldOfStudy: cmd.FieldOfStudy,
		StartDate:   startDate,
		EndDate:     endDate,
		IsCurrent:   cmd.IsCurrent,
		Description: cmd.Description,
		Grade:       cmd.Grade,
		Activities:  cmd.Activities,
		Achievements: cmd.Achievements,
	})
	if err != nil {
		return uuid.Nil, err
	}

	// Save to repository
	if err := s.userEducationRepo.Create.New(ctx, userEducation); err != nil {
		return uuid.Nil, err
	}

	return userEducation.ID(), nil
}
