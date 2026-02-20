package update

import (
	"context"
	"nfxid/modules/clients/domain/apps"
	"nfxid/modules/clients/infrastructure/rdb/models"
	"nfxid/modules/clients/infrastructure/repository/apps/mapper"
)

// Generic 通用更新 App，实现 apps.Update 接口
func (h *Handler) Generic(ctx context.Context, a *apps.App) error {
	m := mapper.AppDomainToModel(a)
	updates := mapper.AppModelToUpdates(m)
	return h.db.WithContext(ctx).
		Model(&models.Application{}).
		Where("id = ?", a.ID()).
		Updates(updates).Error
}
