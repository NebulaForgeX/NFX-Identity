package get

import (
	"context"
	"errors"
	"nfxid/modules/access/domain/permissions"
	"nfxid/modules/access/infrastructure/rdb/models"
	"nfxid/modules/access/infrastructure/repository/permissions/mapper"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// ByID 根据 ID 获取 Permission，实现 permissions.Get 接口
func (h *Handler) ByID(ctx context.Context, id uuid.UUID) (*permissions.Permission, error) {
	var m models.Permission
	if err := h.db.WithContext(ctx).Where("id = ?", id).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, permissions.ErrPermissionNotFound
		}
		return nil, err
	}
	return mapper.PermissionModelToDomain(&m), nil
}
