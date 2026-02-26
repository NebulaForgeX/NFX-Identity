package update

import (
	"context"
	"nfxidentity/modules/directory/domain/user_occupations"
	"nfxidentity/modules/directory/infrastructure/rdb/models"
	"nfxidentity/modules/directory/infrastructure/repository/user_occupations/mapper"
)

// Generic 通用更新 UserOccupation，实现 user_occupations.Update 接口
func (h *Handler) Generic(ctx context.Context, uo *user_occupations.UserOccupation) error {
	m := mapper.UserOccupationDomainToModel(uo)
	updates := mapper.UserOccupationModelToUpdates(m)
	return h.db.WithContext(ctx).
		Model(&models.UserOccupation{}).
		Where("id = ?", uo.ID()).
		Updates(updates).Error
}
