package get

import (
	"context"
	"nfxid/modules/auth/domain/user_role"
	"nfxid/modules/auth/infrastructure/rdb/models"
	"nfxid/modules/auth/infrastructure/repository/mapper"

	"github.com/google/uuid"
)

// ByRoleID 根据 RoleID 获取 UserRole 列表，实现 user_role.Get 接口
func (h *Handler) ByRoleID(ctx context.Context, roleID uuid.UUID) ([]*user_role.UserRole, error) {
	var models []models.UserRole
	if err := h.db.WithContext(ctx).
		Where("role_id = ?", roleID).
		Order("created_at DESC").
		Find(&models).Error; err != nil {
		return nil, err
	}

	entities := make([]*user_role.UserRole, len(models))
	for i := range models {
		entities[i] = mapper.UserRoleModelToDomain(&models[i])
	}
	return entities, nil
}
