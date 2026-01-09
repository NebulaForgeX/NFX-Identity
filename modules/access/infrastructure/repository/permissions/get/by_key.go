package get

import (
	"context"
	"errors"
	"nfxid/modules/access/domain/permissions"
	"nfxid/modules/access/infrastructure/rdb/models"
	"nfxid/modules/access/infrastructure/repository/permissions/mapper"

	"gorm.io/gorm"
)

// ByKey 根据 Key 获取 Permission，实现 permissions.Get 接口
func (h *Handler) ByKey(ctx context.Context, key string) (*permissions.Permission, error) {
	var m models.Permission
	if err := h.db.WithContext(ctx).Where("key = ?", key).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, permissions.ErrPermissionNotFound
		}
		return nil, err
	}
	return mapper.PermissionModelToDomain(&m), nil
}
