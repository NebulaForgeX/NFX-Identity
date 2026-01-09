package create

import (
	"context"
	"nfxid/modules/access/domain/scope_permissions"
	"nfxid/modules/access/infrastructure/repository/scope_permissions/mapper"
)

// New 创建新的 ScopePermission，实现 scope_permissions.Create 接口
func (h *Handler) New(ctx context.Context, sp *scope_permissions.ScopePermission) error {
	m := mapper.ScopePermissionDomainToModel(sp)
	return h.db.WithContext(ctx).Create(&m).Error
}
