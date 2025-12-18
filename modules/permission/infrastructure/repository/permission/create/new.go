package create

import (
	"context"
	permissionDomain "nfxid/modules/permission/domain/permission"
	"nfxid/modules/permission/infrastructure/repository/mapper"
)

// New 创建新的 Permission，实现 permissionDomain.Create 接口
func (h *Handler) New(ctx context.Context, p *permissionDomain.Permission) error {
	m := mapper.PermissionDomainToModel(p)
	return h.db.WithContext(ctx).Create(&m).Error
}
