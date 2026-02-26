package update

import (
	"context"
	"nfxidentity/modules/system/domain/system_state"
	"nfxidentity/modules/system/infrastructure/rdb/models"
	"nfxidentity/modules/system/infrastructure/repository/system_state/mapper"
)

// Generic 通用更新 SystemState，实现 system_state.Update 接口
func (h *Handler) Generic(ctx context.Context, ss *system_state.SystemState) error {
	m := mapper.SystemStateDomainToModel(ss)
	updates := mapper.SystemStateModelToUpdates(m)
	return h.db.WithContext(ctx).
		Model(&models.SystemState{}).
		Where("id = ?", ss.ID()).
		Updates(updates).Error
}
