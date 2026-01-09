package create

import (
	"context"
	"nfxid/modules/directory/domain/user_occupations"
	"nfxid/modules/directory/infrastructure/repository/user_occupations/mapper"
)

// New 创建新的 UserOccupation，实现 user_occupations.Create 接口
func (h *Handler) New(ctx context.Context, uo *user_occupations.UserOccupation) error {
	m := mapper.UserOccupationDomainToModel(uo)
	return h.db.WithContext(ctx).Create(&m).Error
}
