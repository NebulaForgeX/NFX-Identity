package update

import (
	"context"
	"nfxid/modules/auth/domain/profile"
	"nfxid/modules/auth/infrastructure/rdb/models"
	"nfxid/modules/auth/infrastructure/repository/mapper"
)

// Generic 通用更新 Profile，实现 profile.Update 接口
func (h *Handler) Generic(ctx context.Context, p *profile.Profile) error {
	m := mapper.ProfileDomainToModel(p)
	updates := mapper.ProfileModelsToUpdates(m)
	return h.db.WithContext(ctx).
		Model(&models.Profile{}).
		Where("id = ?", p.ID()).
		Updates(updates).Error
}
