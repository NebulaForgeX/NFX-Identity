package get

import (
	"context"
	"errors"
	"nfxid/modules/auth/domain/role"
	roleDomainErrors "nfxid/modules/auth/domain/role/errors"
	"nfxid/modules/auth/infrastructure/rdb/models"
	"nfxid/modules/auth/infrastructure/repository/mapper"

	"gorm.io/gorm"
)

// ByName 根据 Name 获取 Role，实现 role.Get 接口
func (h *Handler) ByName(ctx context.Context, name string) (*role.Role, error) {
	var m models.Role
	if err := h.db.WithContext(ctx).Where("name = ?", name).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, roleDomainErrors.ErrRoleNotFound
		}
		return nil, err
	}
	return mapper.RoleModelToDomain(&m), nil
}
