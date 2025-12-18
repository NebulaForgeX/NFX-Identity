package single

import (
	"context"
	"errors"
	permissionDomainErrors "nfxid/modules/permission/domain/permission/errors"
	permissionDomainViews "nfxid/modules/permission/domain/permission/views"
	"nfxid/modules/permission/infrastructure/rdb/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// ByID 根据 ID 获取 Permission，实现 permissionDomain.Single 接口
func (h *Handler) ByID(ctx context.Context, permissionID uuid.UUID) (*permissionDomainViews.PermissionView, error) {
	var m models.Permission
	if err := h.db.WithContext(ctx).Where("id = ?", permissionID).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, permissionDomainErrors.ErrPermissionNotFound
		}
		return nil, err
	}
	view := permissionModelToDomainView(&m)
	return &view, nil
}
