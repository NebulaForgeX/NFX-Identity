package get

import (
	"context"
	"errors"
	"nfxid/modules/access/domain/roles"
	"nfxid/modules/access/infrastructure/rdb/models"
	"nfxid/modules/access/infrastructure/repository/roles/mapper"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// ByID 根据 ID 获取 Role，实现 roles.Get 接口
func (h *Handler) ByID(ctx context.Context, id uuid.UUID) (*roles.Role, error) {
	var m models.Role
	if err := h.db.WithContext(ctx).Where("id = ?", id).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, roles.ErrRoleNotFound
		}
		return nil, err
	}
	return mapper.RoleModelToDomain(&m), nil
}
