package get

import (
	"context"
	"errors"
	tenantsErr "nfxidentity/errors/src/tenants"
	"nfxidentity/modules/tenants/domain/member_app_roles"
	"nfxidentity/modules/tenants/infrastructure/rdb/models"
	"nfxidentity/modules/tenants/infrastructure/repository/member_app_roles/mapper"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// ByMemberIDAndAppIDAndRoleID 根据 MemberID、AppID 和 RoleID 获取 MemberAppRole，实现 member_app_roles.Get 接口
func (h *Handler) ByMemberIDAndAppIDAndRoleID(ctx context.Context, memberID, appID, roleID uuid.UUID) (*member_app_roles.MemberAppRole, error) {
	var m models.MemberAppRole
	if err := h.db.WithContext(ctx).
		Where("member_id = ? AND application_id = ? AND role_id = ?", memberID, appID, roleID).
		First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, tenantsErr.ErrMemberAppRoleNotFound
		}
		return nil, err
	}
	return mapper.MemberAppRoleModelToDomain(&m), nil
}
