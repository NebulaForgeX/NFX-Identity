package delete

import (
	"context"
	"nfxid/modules/directory/infrastructure/rdb/models"
)

// ByPhone 根据 Phone 删除 UserPhone，实现 user_phones.Delete 接口
func (h *Handler) ByPhone(ctx context.Context, phone string) error {
	return h.db.WithContext(ctx).
		Where("phone = ?", phone).
		Delete(&models.UserPhone{}).Error
}
