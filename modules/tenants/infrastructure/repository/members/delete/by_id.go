package delete

import (
	"context"
	"errors"
	"nfxid/modules/tenants/domain/members"
	"nfxid/modules/tenants/infrastructure/rdb/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// ByID 根据 ID 删除 Member，实现 members.Delete 接口
func (h *Handler) ByID(ctx context.Context, id uuid.UUID) error {
	result := h.db.WithContext(ctx).
		Where("id = ?", id).
		Delete(&models.Member{})

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return members.ErrMemberNotFound
	}
	return nil
}
