package apps

import (
	"context"
	appCommands "nfxid/modules/clients/application/apps/commands"
	appDomain "nfxid/modules/clients/domain/apps"

	"github.com/google/uuid"
)

// CreateApp 创建应用
func (s *Service) CreateApp(ctx context.Context, cmd appCommands.CreateAppCmd) (uuid.UUID, error) {
	// Check if app_id already exists
	if exists, _ := s.appRepo.Check.ByAppID(ctx, cmd.AppID); exists {
		return uuid.Nil, appDomain.ErrAppIDAlreadyExists
	}

	// Create domain entity
	app, err := appDomain.NewApp(appDomain.NewAppParams{
		AppID:       cmd.AppID,
		TenantID:    cmd.TenantID,
		Name:        cmd.Name,
		Description: cmd.Description,
		Type:        cmd.Type,
		Status:      cmd.Status,
		Environment: cmd.Environment,
		CreatedBy:   cmd.CreatedBy,
		Metadata:    cmd.Metadata,
	})
	if err != nil {
		return uuid.Nil, err
	}

	// Save to repository
	if err := s.appRepo.Create.New(ctx, app); err != nil {
		return uuid.Nil, err
	}

	return app.ID(), nil
}
