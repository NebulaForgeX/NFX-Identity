package create

import (
	"context"
	"nfxid/modules/access/domain/permissions"
	"nfxid/modules/access/infrastructure/repository/permissions/mapper"
)

// New 创建新的 Permission，实现 permissions.Create 接口
func (h *Handler) New(ctx context.Context, p *permissions.Permission) error {
	m := mapper.PermissionDomainToModel(p)
	return h.db.WithContext(ctx).Create(&m).Error
}
