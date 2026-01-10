package apps

import (
	"context"
	appCommands "nfxid/modules/clients/application/apps/commands"
)

// UpdateApp 更新应用
func (s *Service) UpdateApp(ctx context.Context, cmd appCommands.UpdateAppCmd) error {
	// Get domain entity
	app, err := s.appRepo.Get.ByID(ctx, cmd.AppID)
	if err != nil {
		return err
	}

	// Update domain entity
	if err := app.Update(cmd.Name, cmd.Description, cmd.Type, cmd.Environment, cmd.Metadata, cmd.UpdatedBy); err != nil {
		return err
	}

	// Save to repository
	return s.appRepo.Update.Generic(ctx, app)
}

// UpdateAppStatus 更新应用状态
func (s *Service) UpdateAppStatus(ctx context.Context, cmd appCommands.UpdateAppStatusCmd) error {
	// Get domain entity
	app, err := s.appRepo.Get.ByID(ctx, cmd.AppID)
	if err != nil {
		return err
	}

	// Update status domain entity
	if err := app.UpdateStatus(cmd.Status, nil); err != nil {
		return err
	}

	// Save to repository
	return s.appRepo.Update.Status(ctx, cmd.AppID, cmd.Status)
}
