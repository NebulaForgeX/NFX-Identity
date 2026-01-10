package system_state

import (
	"context"
	systemStateResult "nfxid/modules/system/application/system_state/results"

	"github.com/google/uuid"
)

// GetSystemState 根据ID获取系统状态
func (s *Service) GetSystemState(ctx context.Context, systemStateID uuid.UUID) (systemStateResult.SystemStateRO, error) {
	domainEntity, err := s.systemStateRepo.Get.ByID(ctx, systemStateID)
	if err != nil {
		return systemStateResult.SystemStateRO{}, err
	}
	return systemStateResult.SystemStateMapper(domainEntity), nil
}

// GetLatestSystemState 获取最新的系统状态
func (s *Service) GetLatestSystemState(ctx context.Context) (systemStateResult.SystemStateRO, error) {
	domainEntity, err := s.systemStateRepo.Get.Latest(ctx)
	if err != nil {
		return systemStateResult.SystemStateRO{}, err
	}
	return systemStateResult.SystemStateMapper(domainEntity), nil
}
