package get

import (
	"context"
	"errors"
	"nfxid/modules/access/domain/role_permissions"
	"nfxid/modules/access/infrastructure/rdb/models"
	"nfxid/modules/access/infrastructure/repository/role_permissions/mapper"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// ByRoleIDAndPermissionID 根据 RoleID 和 PermissionID 获取 RolePermission，实现 role_permissions.Get 接口
func (h *Handler) ByRoleIDAndPermissionID(ctx context.Context, roleID, permissionID uuid.UUID) (*role_permissions.RolePermission, error) {
	var m models.RolePermission
	if err := h.db.WithContext(ctx).
		Where("role_id = ? AND permission_id = ?", roleID, permissionID).
		First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, role_permissions.ErrRolePermissionNotFound
		}
		return nil, err
	}
	return mapper.RolePermissionModelToDomain(&m), nil
}
