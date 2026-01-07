package update

import (
	"context"
	"nfxid/modules/auth/domain/profile"
	"nfxid/modules/auth/infrastructure/rdb/models"
	"nfxid/modules/auth/infrastructure/repository/mapper"
)

// AndInsert Upsert 操作：基于 user_id，如果存在则更新，不存在则插入，实现 profile.Update 接口
func (h *Handler) AndInsert(ctx context.Context, p *profile.Profile) error {
	m := mapper.ProfileDomainToModel(p)
	
	// 尝试查找现有记录
	var existing models.Profile
	err := h.db.WithContext(ctx).
		Where("user_id = ?", p.UserID()).
		First(&existing).Error
	
	if err != nil {
		// 如果不存在，创建新记录
		if err.Error() == "record not found" {
			return h.db.WithContext(ctx).Create(m).Error
		}
		return err
	}
	
	// 如果存在，更新记录
	updates := mapper.ProfileModelsToUpdates(m)
	return h.db.WithContext(ctx).
		Model(&models.Profile{}).
		Where("user_id = ?", p.UserID()).
		Updates(updates).Error
}
