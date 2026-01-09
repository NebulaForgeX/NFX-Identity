package create

import (
	"context"
	"nfxid/modules/tenants/domain/groups"
	"nfxid/modules/tenants/infrastructure/repository/groups/mapper"
)

// New 创建新的 Group，实现 groups.Create 接口
func (h *Handler) New(ctx context.Context, g *groups.Group) error {
	m := mapper.GroupDomainToModel(g)
	return h.db.WithContext(ctx).Create(&m).Error
}
