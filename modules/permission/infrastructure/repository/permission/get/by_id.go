package get

import (
	"context"
	"errors"
	permissionDomain "nfxid/modules/permission/domain/permission"
	permissionDomainErrors "nfxid/modules/permission/domain/permission/errors"
	"nfxid/modules/permission/infrastructure/rdb/models"
	"nfxid/modules/permission/infrastructure/repository/mapper"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// ByID 根据 ID 获取 Permission，实现 permissionDomain.Get 接口
func (h *Handler) ByID(ctx context.Context, id uuid.UUID) (*permissionDomain.Permission, error) {
	var m models.Permission
	if err := h.db.WithContext(ctx).Where("id = ?", id).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, permissionDomainErrors.ErrPermissionNotFound
		}
		return nil, err
	}
	return mapper.PermissionModelToDomain(&m), nil
}
