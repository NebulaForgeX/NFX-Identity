package list

import (
	"context"
	userRoleDomainViews "nfxid/modules/auth/domain/user_role/views"
	"nfxid/modules/auth/infrastructure/query/mapper"
	"nfxid/modules/auth/infrastructure/rdb/models"
	"nfxid/pkgs/utils/slice"

	"github.com/google/uuid"
)

// ByRoleID 根据 RoleID 获取 UserRole 列表，实现 userRoleDomain.List 接口
func (h *Handler) ByRoleID(ctx context.Context, roleID uuid.UUID) ([]*userRoleDomainViews.UserRoleView, error) {
	var items []models.UserRole
	if err := h.db.WithContext(ctx).
		Where("role_id = ?", roleID).
		Order("created_at DESC").
		Find(&items).Error; err != nil {
		return nil, err
	}
	result := slice.MapP(items, mapper.UserRoleModelToDomain)
	// Convert to pointers
	pointerResult := make([]*userRoleDomainViews.UserRoleView, len(result))
	for i := range result {
		pointerResult[i] = &result[i]
	}
	return pointerResult, nil
}
