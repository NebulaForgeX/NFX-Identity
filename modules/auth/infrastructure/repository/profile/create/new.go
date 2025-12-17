package create

import (
	"context"
	"nfxid/modules/auth/domain/profile"
	"nfxid/modules/auth/infrastructure/repository/mapper"
)

// New 创建新的 Profile，实现 profile.Create 接口
func (h *Handler) New(ctx context.Context, p *profile.Profile) error {
	m := mapper.ProfileDomainToModel(p)
	return h.db.WithContext(ctx).Create(&m).Error
}
