package create

import (
	"context"
	"nfxid/modules/clients/domain/apps"
	"nfxid/modules/clients/infrastructure/repository/apps/mapper"
)

// New 创建新的 App，实现 apps.Create 接口
func (h *Handler) New(ctx context.Context, a *apps.App) error {
	m := mapper.AppDomainToModel(a)
	return h.db.WithContext(ctx).Create(&m).Error
}
