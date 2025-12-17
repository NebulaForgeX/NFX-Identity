package update

import (
	"context"
	occupation "nfxid/modules/auth/domain/profile_occupation"
	"nfxid/modules/auth/infrastructure/rdb/models"
	"nfxid/modules/auth/infrastructure/repository/mapper"
)

// Generic 通用更新 Occupation，实现 occupation.Update 接口
func (h *Handler) Generic(ctx context.Context, o *occupation.Occupation) error {
	m := mapper.OccupationDomainToModel(o)
	updates := mapper.OccupationModelsToUpdates(m)
	return h.db.WithContext(ctx).
		Model(&models.Occupation{}).
		Where("id = ?", o.ID()).
		Updates(updates).Error
}
