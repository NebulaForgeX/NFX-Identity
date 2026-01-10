package user_educations

import (
	"context"
	"time"
	userEducationCommands "nfxid/modules/directory/application/user_educations/commands"
)

// UpdateUserEducation 更新用户教育经历
func (s *Service) UpdateUserEducation(ctx context.Context, cmd userEducationCommands.UpdateUserEducationCmd) error {
	// Get domain entity
	userEducation, err := s.userEducationRepo.Get.ByID(ctx, cmd.UserEducationID)
	if err != nil {
		return err
	}

	var startDate, endDate *time.Time
	if cmd.StartDate != nil && *cmd.StartDate != "" {
		parsed, err := time.Parse(time.RFC3339, *cmd.StartDate)
		if err != nil {
			return err
		}
		startDate = &parsed
	}
	if cmd.EndDate != nil && *cmd.EndDate != "" {
		parsed, err := time.Parse(time.RFC3339, *cmd.EndDate)
		if err != nil {
			return err
		}
		endDate = &parsed
	}

	// Update domain entity
	if err := userEducation.Update(cmd.School, cmd.Degree, cmd.Major, cmd.FieldOfStudy, startDate, endDate, cmd.IsCurrent, cmd.Description, cmd.Grade, cmd.Activities, cmd.Achievements); err != nil {
		return err
	}

	// Save to repository
	return s.userEducationRepo.Update.Generic(ctx, userEducation)
}
