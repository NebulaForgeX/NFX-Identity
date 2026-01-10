package user_educations

import (
	"context"
	userEducationCommands "nfxid/modules/directory/application/user_educations/commands"
)

// DeleteUserEducation 删除用户教育经历（软删除）
func (s *Service) DeleteUserEducation(ctx context.Context, cmd userEducationCommands.DeleteUserEducationCmd) error {
	// Get domain entity
	userEducation, err := s.userEducationRepo.Get.ByID(ctx, cmd.UserEducationID)
	if err != nil {
		return err
	}

	// Delete domain entity (soft delete)
	if err := userEducation.Delete(); err != nil {
		return err
	}

	// Save to repository
	return s.userEducationRepo.Update.Generic(ctx, userEducation)
}
