package get

import (
	"context"
	"errors"
	"nfxid/modules/access/domain/roles"
	"nfxid/modules/access/infrastructure/rdb/models"
	"nfxid/modules/access/infrastructure/repository/roles/mapper"

	"gorm.io/gorm"
)

// ByKey 根据 Key 获取 Role，实现 roles.Get 接口
func (h *Handler) ByKey(ctx context.Context, key string) (*roles.Role, error) {
	var m models.Role
	if err := h.db.WithContext(ctx).Where("key = ?", key).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, roles.ErrRoleNotFound
		}
		return nil, err
	}
	return mapper.RoleModelToDomain(&m), nil
}
