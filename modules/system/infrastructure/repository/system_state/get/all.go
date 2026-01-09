package get

import (
	"context"
	"errors"
	"nfxid/modules/system/domain/system_state"
	"nfxid/modules/system/infrastructure/rdb/models"
	"nfxid/modules/system/infrastructure/repository/system_state/mapper"

	"gorm.io/gorm"
)

// All 获取所有 SystemState，实现 system_state.Get 接口
func (h *Handler) All(ctx context.Context) ([]*system_state.SystemState, error) {
	var ms []models.SystemState
	if err := h.db.WithContext(ctx).Order("created_at DESC").Find(&ms).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []*system_state.SystemState{}, nil
		}
		return nil, err
	}
	
	result := make([]*system_state.SystemState, len(ms))
	for i := range ms {
		result[i] = mapper.SystemStateModelToDomain(&ms[i])
	}
	return result, nil
}
