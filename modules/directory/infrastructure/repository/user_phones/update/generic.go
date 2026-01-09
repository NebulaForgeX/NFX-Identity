package update

import (
	"context"
	"nfxid/modules/directory/domain/user_phones"
	"nfxid/modules/directory/infrastructure/rdb/models"
	"nfxid/modules/directory/infrastructure/repository/user_phones/mapper"
)

// Generic 通用更新 UserPhone，实现 user_phones.Update 接口
func (h *Handler) Generic(ctx context.Context, up *user_phones.UserPhone) error {
	m := mapper.UserPhoneDomainToModel(up)
	updates := mapper.UserPhoneModelToUpdates(m)
	return h.db.WithContext(ctx).
		Model(&models.UserPhone{}).
		Where("id = ?", up.ID()).
		Updates(updates).Error
}
