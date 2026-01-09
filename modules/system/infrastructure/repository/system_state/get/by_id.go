package get

import (
	"context"
	"errors"
	"nfxid/modules/system/domain/system_state"
	"nfxid/modules/system/infrastructure/rdb/models"
	"nfxid/modules/system/infrastructure/repository/system_state/mapper"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// ByID 根据 ID 获取 SystemState，实现 system_state.Get 接口
func (h *Handler) ByID(ctx context.Context, id uuid.UUID) (*system_state.SystemState, error) {
	var m models.SystemState
	if err := h.db.WithContext(ctx).Where("id = ?", id).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, system_state.ErrSystemStateNotFound
		}
		return nil, err
	}
	return mapper.SystemStateModelToDomain(&m), nil
}
