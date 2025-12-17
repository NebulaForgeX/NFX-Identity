package create

import (
	"context"
	occupation "nfxid/modules/auth/domain/profile_occupation"
	"nfxid/modules/auth/infrastructure/repository/mapper"
)

// New 创建新的 Occupation，实现 occupation.Create 接口
func (h *Handler) New(ctx context.Context, o *occupation.Occupation) error {
	m := mapper.OccupationDomainToModel(o)
	return h.db.WithContext(ctx).Create(&m).Error
}
