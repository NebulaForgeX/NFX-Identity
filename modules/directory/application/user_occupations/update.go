package user_occupations

import (
	"context"
	"time"
	userOccupationCommands "nfxid/modules/directory/application/user_occupations/commands"
)

// UpdateUserOccupation 更新用户职业经历
func (s *Service) UpdateUserOccupation(ctx context.Context, cmd userOccupationCommands.UpdateUserOccupationCmd) error {
	// Get domain entity
	userOccupation, err := s.userOccupationRepo.Get.ByID(ctx, cmd.UserOccupationID)
	if err != nil {
		return err
	}

	var startDate, endDate *time.Time
	if cmd.StartDate != nil && *cmd.StartDate != "" {
		parsed, err := parseDateString(*cmd.StartDate)
		if err != nil {
			return err
		}
		startDate = parsed
	}
	if cmd.EndDate != nil && *cmd.EndDate != "" {
		parsed, err := parseDateString(*cmd.EndDate)
		if err != nil {
			return err
		}
		endDate = parsed
	}

	// Update domain entity
	if err := userOccupation.Update(cmd.Company, cmd.Position, cmd.Department, cmd.Industry, cmd.Location, cmd.EmploymentType, startDate, endDate, cmd.IsCurrent, cmd.Description, cmd.Responsibilities, cmd.Achievements, cmd.SkillsUsed); err != nil {
		return err
	}

	// Save to repository
	return s.userOccupationRepo.Update.Generic(ctx, userOccupation)
}
