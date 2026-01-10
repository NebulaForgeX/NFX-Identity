package apps

import (
	"context"
	appCommands "nfxid/modules/clients/application/apps/commands"
)

// DeleteApp 删除应用（软删除）
func (s *Service) DeleteApp(ctx context.Context, cmd appCommands.DeleteAppCmd) error {
	// Get domain entity
	app, err := s.appRepo.Get.ByID(ctx, cmd.AppID)
	if err != nil {
		return err
	}

	// Delete domain entity (soft delete)
	if err := app.Delete(); err != nil {
		return err
	}

	// Save to repository
	return s.appRepo.Update.Generic(ctx, app)
}
