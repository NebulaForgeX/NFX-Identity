package get

import (
	"context"
	"nfxidentity/modules/tenants/domain/member_roles"
	"nfxidentity/modules/tenants/infrastructure/rdb/models"
	"nfxidentity/modules/tenants/infrastructure/repository/member_roles/mapper"

	"github.com/google/uuid"
)

// ByRoleID 根据 RoleID 获取 MemberRole 列表，实现 member_roles.Get 接口
func (h *Handler) ByRoleID(ctx context.Context, roleID uuid.UUID) ([]*member_roles.MemberRole, error) {
	var ms []models.MemberRole
	if err := h.db.WithContext(ctx).
		Where("role_id = ?", roleID).
		Find(&ms).Error; err != nil {
		return nil, err
	}

	result := make([]*member_roles.MemberRole, len(ms))
	for i, m := range ms {
		result[i] = mapper.MemberRoleModelToDomain(&m)
	}
	return result, nil
}
