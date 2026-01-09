package update

import (
	"context"
	"nfxid/modules/directory/domain/user_educations"
	"nfxid/modules/directory/infrastructure/rdb/models"
	"nfxid/modules/directory/infrastructure/repository/user_educations/mapper"
)

// Generic 通用更新 UserEducation，实现 user_educations.Update 接口
func (h *Handler) Generic(ctx context.Context, ue *user_educations.UserEducation) error {
	m := mapper.UserEducationDomainToModel(ue)
	updates := mapper.UserEducationModelToUpdates(m)
	return h.db.WithContext(ctx).
		Model(&models.UserEducation{}).
		Where("id = ?", ue.ID()).
		Updates(updates).Error
}
