package check

import (
	"context"
	"nfxid/modules/tenants/infrastructure/rdb/models"

	"github.com/google/uuid"
)

// ByTenantIDAndMemberIDAndRoleID 根据 TenantID、MemberID 和 RoleID 检查 MemberRole 是否存在，实现 member_roles.Check 接口
func (h *Handler) ByTenantIDAndMemberIDAndRoleID(ctx context.Context, tenantID, memberID, roleID uuid.UUID) (bool, error) {
	var count int64
	err := h.db.WithContext(ctx).
		Model(&models.MemberRole{}).
		Where("tenant_id = ? AND member_id = ? AND role_id = ?", tenantID, memberID, roleID).
		Count(&count).Error
	return count > 0, err
}
