package system_state

import (
	"context"
	"errors"
	"time"

	systemStateResult "nfxid/modules/system/application/system_state/results"
	systemStateDomain "nfxid/modules/system/domain/system_state"

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
// 如果系统未初始化（没有记录），返回默认的未初始化状态
func (s *Service) GetLatestSystemState(ctx context.Context) (systemStateResult.SystemStateRO, error) {
	domainEntity, err := s.systemStateRepo.Get.Latest(ctx)
	if err != nil {
		// 如果是未找到记录的错误，返回默认的未初始化状态
		if errors.Is(err, systemStateDomain.ErrSystemStateNotFound) {
			now := time.Now()
			return systemStateResult.SystemStateRO{
				ID:                   uuid.Nil,
				Initialized:          false,
				InitializedAt:        nil,
				InitializationVersion: nil,
				LastResetAt:          nil,
				LastResetBy:          nil,
				ResetCount:           0,
				Metadata:             make(map[string]interface{}),
				CreatedAt:            now,
				UpdatedAt:            now,
			}, nil
		}
		return systemStateResult.SystemStateRO{}, err
	}
	return systemStateResult.SystemStateMapper(domainEntity), nil
}
