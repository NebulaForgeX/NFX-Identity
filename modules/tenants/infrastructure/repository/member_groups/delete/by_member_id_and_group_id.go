package delete

import (
	"context"
	tenantsErr "nfxidentity/errors/src/tenants"
	"nfxidentity/modules/tenants/infrastructure/rdb/models"

	"github.com/google/uuid"
)

// ByMemberIDAndGroupID 根据 MemberID 和 GroupID 删除 MemberGroup，实现 member_groups.Delete 接口
func (h *Handler) ByMemberIDAndGroupID(ctx context.Context, memberID, groupID uuid.UUID) error {
	result := h.db.WithContext(ctx).
		Where("member_id = ? AND group_id = ?", memberID, groupID).
		Delete(&models.MemberGroup{})

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return tenantsErr.ErrMemberGroupNotFound
	}
	return nil
}
