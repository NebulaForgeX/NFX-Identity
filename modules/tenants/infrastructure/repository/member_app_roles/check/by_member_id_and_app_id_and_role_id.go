package check

import (
	"context"
	"nfxid/modules/tenants/infrastructure/rdb/models"

	"github.com/google/uuid"
)

// ByMemberIDAndAppIDAndRoleID 根据 MemberID、AppID 和 RoleID 检查 MemberAppRole 是否存在，实现 member_app_roles.Check 接口
func (h *Handler) ByMemberIDAndAppIDAndRoleID(ctx context.Context, memberID, appID, roleID uuid.UUID) (bool, error) {
	var count int64
	err := h.db.WithContext(ctx).
		Model(&models.MemberAppRole{}).
		Where("member_id = ? AND app_id = ? AND role_id = ?", memberID, appID, roleID).
		Count(&count).Error
	return count > 0, err
}
