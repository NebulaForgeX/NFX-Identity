package delete

import (
	"context"
	"nfxid/modules/tenants/domain/member_groups"
	"nfxid/modules/tenants/infrastructure/rdb/models"

	"github.com/google/uuid"
)

// ByID 根据 ID 删除 MemberGroup，实现 member_groups.Delete 接口
func (h *Handler) ByID(ctx context.Context, id uuid.UUID) error {
	result := h.db.WithContext(ctx).
		Where("id = ?", id).
		Delete(&models.MemberGroup{})

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return member_groups.ErrMemberGroupNotFound
	}
	return nil
}
