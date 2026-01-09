package create

import (
	"context"
	"nfxid/modules/directory/domain/user_phones"
	"nfxid/modules/directory/infrastructure/repository/user_phones/mapper"
)

// New 创建新的 UserPhone，实现 user_phones.Create 接口
func (h *Handler) New(ctx context.Context, up *user_phones.UserPhone) error {
	m := mapper.UserPhoneDomainToModel(up)
	return h.db.WithContext(ctx).Create(&m).Error
}
