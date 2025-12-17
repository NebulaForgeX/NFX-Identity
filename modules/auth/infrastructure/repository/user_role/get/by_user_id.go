package get

import (
	"context"
	"nfxid/modules/auth/domain/user_role"
	"nfxid/modules/auth/infrastructure/rdb/models"
	"nfxid/modules/auth/infrastructure/repository/mapper"

	"github.com/google/uuid"
)

// ByUserID 根据 UserID 获取 UserRole 列表，实现 user_role.Get 接口
func (h *Handler) ByUserID(ctx context.Context, userID uuid.UUID) ([]*user_role.UserRole, error) {
	var models []models.UserRole
	if err := h.db.WithContext(ctx).
		Where("user_id = ?", userID).
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
