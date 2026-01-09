package check

import (
	"context"
	"nfxid/modules/tenants/infrastructure/rdb/models"

	"github.com/google/uuid"
)

// ByMemberID 根据 MemberID 检查 Member 是否存在，实现 members.Check 接口
func (h *Handler) ByMemberID(ctx context.Context, memberID uuid.UUID) (bool, error) {
	var count int64
	err := h.db.WithContext(ctx).
		Model(&models.Member{}).
		Where("member_id = ?", memberID).
		Count(&count).Error
	return count > 0, err
}
