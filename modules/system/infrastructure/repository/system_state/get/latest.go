package get

import (
	"context"
	"errors"
	"nfxid/modules/system/domain/system_state"
	"nfxid/modules/system/infrastructure/rdb/models"
	"nfxid/modules/system/infrastructure/repository/system_state/mapper"

	"gorm.io/gorm"
)

// Latest 获取最新的 SystemState，实现 system_state.Get 接口
func (h *Handler) Latest(ctx context.Context) (*system_state.SystemState, error) {
	var m models.SystemState
	if err := h.db.WithContext(ctx).Order("created_at DESC").First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, system_state.ErrSystemStateNotFound
		}
		return nil, err
	}
	return mapper.SystemStateModelToDomain(&m), nil
}
