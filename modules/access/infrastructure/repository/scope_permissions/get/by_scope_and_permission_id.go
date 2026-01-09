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

// ByScopeAndPermissionID 根据 Scope 和 PermissionID 获取 ScopePermission，实现 scope_permissions.Get 接口
func (h *Handler) ByScopeAndPermissionID(ctx context.Context, scope string, permissionID uuid.UUID) (*scope_permissions.ScopePermission, error) {
	var m models.ScopePermission
	if err := h.db.WithContext(ctx).
		Where("scope = ? AND permission_id = ?", scope, permissionID).
		First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, scope_permissions.ErrScopePermissionNotFound
		}
		return nil, err
	}
	return mapper.ScopePermissionModelToDomain(&m), nil
}
