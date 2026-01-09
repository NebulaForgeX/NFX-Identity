package update

import (
	"context"
	"time"
	"nfxid/modules/directory/infrastructure/rdb/models"

	"github.com/google/uuid"
)

// SetPrimary 设置为主手机号，实现 user_phones.Update 接口
func (h *Handler) SetPrimary(ctx context.Context, id uuid.UUID) error {
	// 先获取这个手机号的 user_id
	var m models.UserPhone
	if err := h.db.WithContext(ctx).
		Where("id = ?", id).
		First(&m).Error; err != nil {
		return err
	}

	// 先将该用户的所有手机号设置为非主手机号
	updates1 := map[string]any{
		models.UserPhoneCols.IsPrimary: false,
		models.UserPhoneCols.UpdatedAt: time.Now().UTC(),
	}
	if err := h.db.WithContext(ctx).
		Model(&models.UserPhone{}).
		Where("user_id = ?", m.UserID).
		Updates(updates1).Error; err != nil {
		return err
	}

	// 然后设置这个手机号为主手机号
	updates2 := map[string]any{
		models.UserPhoneCols.IsPrimary: true,
		models.UserPhoneCols.UpdatedAt: time.Now().UTC(),
	}
	return h.db.WithContext(ctx).
		Model(&models.UserPhone{}).
		Where("id = ?", id).
		Updates(updates2).Error
}
