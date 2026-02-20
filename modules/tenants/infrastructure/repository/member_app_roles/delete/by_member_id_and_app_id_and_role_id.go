package delete

import (
	"context"
	"nfxid/modules/tenants/infrastructure/rdb/models"

	"github.com/google/uuid"
)

// ByMemberIDAndAppIDAndRoleID 根据 MemberID、AppID 和 RoleID 删除 MemberAppRole，实现 member_app_roles.Delete 接口
func (h *Handler) ByMemberIDAndAppIDAndRoleID(ctx context.Context, memberID, appID, roleID uuid.UUID) error {
	return h.db.WithContext(ctx).
		Where("member_id = ? AND application_id = ? AND role_id = ?", memberID, appID, roleID).
		Delete(&models.MemberAppRole{}).Error
}
