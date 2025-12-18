package single

import (
	"context"
	"errors"
	permissionDomainErrors "nfxid/modules/permission/domain/permission/errors"
	permissionDomainViews "nfxid/modules/permission/domain/permission/views"
	"nfxid/modules/permission/infrastructure/rdb/models"

	"gorm.io/gorm"
)

// ByTag 根据 Tag 获取 Permission，实现 permissionDomain.Single 接口
func (h *Handler) ByTag(ctx context.Context, tag string) (*permissionDomainViews.PermissionView, error) {
	var m models.Permission
	if err := h.db.WithContext(ctx).Where("tag = ?", tag).Where("deleted_at IS NULL").First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, permissionDomainErrors.ErrPermissionNotFound
		}
		return nil, err
	}
	view := permissionModelToDomainView(&m)
	return &view, nil
}
