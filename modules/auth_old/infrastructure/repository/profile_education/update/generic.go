package update

import (
	"context"
	education "nfxid/modules/auth/domain/profile_education"
	"nfxid/modules/auth/infrastructure/rdb/models"
	"nfxid/modules/auth/infrastructure/repository/mapper"
)

// Generic 通用更新 Education，实现 education.Update 接口
func (h *Handler) Generic(ctx context.Context, e *education.Education) error {
	m := mapper.EducationDomainToModel(e)
	updates := mapper.EducationModelsToUpdates(m)
	return h.db.WithContext(ctx).
		Model(&models.Education{}).
		Where("id = ?", e.ID()).
		Updates(updates).Error
}
