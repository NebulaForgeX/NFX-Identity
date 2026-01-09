package create

import (
	"context"
	"nfxid/modules/access/domain/role_permissions"
	"nfxid/modules/access/infrastructure/repository/role_permissions/mapper"
)

// New 创建新的 RolePermission，实现 role_permissions.Create 接口
func (h *Handler) New(ctx context.Context, rp *role_permissions.RolePermission) error {
	m := mapper.RolePermissionDomainToModel(rp)
	return h.db.WithContext(ctx).Create(&m).Error
}
