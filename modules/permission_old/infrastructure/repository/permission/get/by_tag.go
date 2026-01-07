package get

import (
	"context"
	"errors"
	permissionDomain "nfxid/modules/permission/domain/permission"
	permissionDomainErrors "nfxid/modules/permission/domain/permission/errors"
	"nfxid/modules/permission/infrastructure/rdb/models"
	"nfxid/modules/permission/infrastructure/repository/mapper"

	"gorm.io/gorm"
)

// ByTag 根据 Tag 获取 Permission，实现 permissionDomain.Get 接口
func (h *Handler) ByTag(ctx context.Context, tag string) (*permissionDomain.Permission, error) {
	var m models.Permission
	if err := h.db.WithContext(ctx).Where("tag = ?", tag).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, permissionDomainErrors.ErrPermissionNotFound
		}
		return nil, err
	}
	return mapper.PermissionModelToDomain(&m), nil
}
