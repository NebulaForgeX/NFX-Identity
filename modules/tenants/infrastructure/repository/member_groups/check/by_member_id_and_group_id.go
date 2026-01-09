package check

import (
	"context"
	"nfxid/modules/tenants/infrastructure/rdb/models"

	"github.com/google/uuid"
)

// ByMemberIDAndGroupID 根据 MemberID 和 GroupID 检查 MemberGroup 是否存在，实现 member_groups.Check 接口
func (h *Handler) ByMemberIDAndGroupID(ctx context.Context, memberID, groupID uuid.UUID) (bool, error) {
	var count int64
	err := h.db.WithContext(ctx).
		Model(&models.MemberGroup{}).
		Where("member_id = ? AND group_id = ?", memberID, groupID).
		Count(&count).Error
	return count > 0, err
}
