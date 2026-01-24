package user_occupations

import (
	"context"
	"time"
	userOccupationCommands "nfxid/modules/directory/application/user_occupations/commands"
	userOccupationDomain "nfxid/modules/directory/domain/user_occupations"

	"github.com/google/uuid"
)

// CreateUserOccupation 创建用户职业经历
func (s *Service) CreateUserOccupation(ctx context.Context, cmd userOccupationCommands.CreateUserOccupationCmd) (uuid.UUID, error) {
	var startDate, endDate *time.Time
	var err error
	if cmd.StartDate != nil && *cmd.StartDate != "" {
		startDate, err = parseDateString(*cmd.StartDate)
		if err != nil {
			return uuid.Nil, err
		}
	}
	if cmd.EndDate != nil && *cmd.EndDate != "" {
		endDate, err = parseDateString(*cmd.EndDate)
		if err != nil {
			return uuid.Nil, err
		}
	}

	// Create domain entity
	userOccupation, err := userOccupationDomain.NewUserOccupation(userOccupationDomain.NewUserOccupationParams{
		UserID:          cmd.UserID,
		Company:         cmd.Company,
		Position:        cmd.Position,
		Department:      cmd.Department,
		Industry:        cmd.Industry,
		Location:        cmd.Location,
		EmploymentType:  cmd.EmploymentType,
		StartDate:       startDate,
		EndDate:         endDate,
		IsCurrent:       cmd.IsCurrent,
		Description:     cmd.Description,
		Responsibilities: cmd.Responsibilities,
		Achievements:    cmd.Achievements,
		SkillsUsed:      cmd.SkillsUsed,
	})
	if err != nil {
		return uuid.Nil, err
	}

	// Save to repository
	if err := s.userOccupationRepo.Create.New(ctx, userOccupation); err != nil {
		return uuid.Nil, err
	}

	return userOccupation.ID(), nil
}
