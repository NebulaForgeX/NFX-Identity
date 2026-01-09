package check

import (
	"context"
	"nfxid/modules/directory/infrastructure/rdb/models"
)

// ByPhone 根据 Phone 检查 UserPhone 是否存在，实现 user_phones.Check 接口
func (h *Handler) ByPhone(ctx context.Context, phone string) (bool, error) {
	var count int64
	err := h.db.WithContext(ctx).
		Model(&models.UserPhone{}).
		Where("phone = ?", phone).
		Count(&count).Error
	return count > 0, err
}
