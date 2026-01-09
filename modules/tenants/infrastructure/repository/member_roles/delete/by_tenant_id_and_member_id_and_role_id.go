package delete

import (
	"context"
	"nfxid/modules/tenants/domain/member_roles"
	"nfxid/modules/tenants/infrastructure/rdb/models"

	"github.com/google/uuid"
)

// ByTenantIDAndMemberIDAndRoleID 根据 TenantID、MemberID 和 RoleID 删除 MemberRole，实现 member_roles.Delete 接口
func (h *Handler) ByTenantIDAndMemberIDAndRoleID(ctx context.Context, tenantID, memberID, roleID uuid.UUID) error {
	result := h.db.WithContext(ctx).
		Where("tenant_id = ? AND member_id = ? AND role_id = ?", tenantID, memberID, roleID).
		Delete(&models.MemberRole{})

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return member_roles.ErrMemberRoleNotFound
	}
	return nil
}
