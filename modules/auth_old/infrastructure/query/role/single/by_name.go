package single

import (
	"context"
	"errors"
	roleDomainErrors "nfxid/modules/auth/domain/role/errors"
	roleDomainViews "nfxid/modules/auth/domain/role/views"
	"nfxid/modules/auth/infrastructure/query/mapper"
	"nfxid/modules/auth/infrastructure/rdb/models"

	"gorm.io/gorm"
)

// ByName 根据名称获取 Role，实现 roleDomain.Single 接口
func (h *Handler) ByName(ctx context.Context, name string) (*roleDomainViews.RoleView, error) {
	var m models.Role
	if err := h.db.WithContext(ctx).Where("name = ?", name).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, roleDomainErrors.ErrRoleViewNotFound
		}
		return nil, err
	}
	view := mapper.RoleModelToDomain(&m)
	return &view, nil
}
