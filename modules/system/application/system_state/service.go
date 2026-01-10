package system_state

import (
	"context"
	systemStateCommands "nfxid/modules/system/application/system_state/commands"
	systemStateDomain "nfxid/modules/system/domain/system_state"
)

type Service struct {
	systemStateRepo *systemStateDomain.Repo
}

func NewService(
	systemStateRepo *systemStateDomain.Repo,
) *Service {
	return &Service{
		systemStateRepo: systemStateRepo,
	}
}

// InitializeSystem 初始化系统
func (s *Service) InitializeSystem(ctx context.Context, cmd systemStateCommands.InitializeSystemCmd) error {
	return s.systemStateRepo.Update.Initialize(ctx, cmd.Version)
}

// ResetSystem 重置系统
func (s *Service) ResetSystem(ctx context.Context, cmd systemStateCommands.ResetSystemCmd) error {
	return s.systemStateRepo.Update.Reset(ctx, cmd.ResetBy)
}

// DeleteSystemState 删除系统状态
func (s *Service) DeleteSystemState(ctx context.Context, cmd systemStateCommands.DeleteSystemStateCmd) error {
	return s.systemStateRepo.Delete.ByID(ctx, cmd.SystemStateID)
}
