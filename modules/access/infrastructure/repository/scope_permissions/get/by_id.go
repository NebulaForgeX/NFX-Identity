package get

import (
	"context"
	"errors"
	"nfxid/modules/access/domain/scope_permissions"
	"nfxid/modules/access/infrastructure/rdb/models"
	"nfxid/modules/access/infrastructure/repository/scope_permissions/mapper"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// ByID 根据 ID 获取 ScopePermission，实现 scope_permissions.Get 接口
func (h *Handler) ByID(ctx context.Context, id uuid.UUID) (*scope_permissions.ScopePermission, error) {
	var m models.ScopePermission
	if err := h.db.WithContext(ctx).Where("id = ?", id).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, scope_permissions.ErrScopePermissionNotFound
		}
		return nil, err
	}
	return mapper.ScopePermissionModelToDomain(&m), nil
}
