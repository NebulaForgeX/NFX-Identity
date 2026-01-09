package create

import (
	"context"
	"nfxid/modules/access/domain/scopes"
	"nfxid/modules/access/infrastructure/repository/scopes/mapper"
)

// New 创建新的 Scope，实现 scopes.Create 接口
func (h *Handler) New(ctx context.Context, s *scopes.Scope) error {
	m := mapper.ScopeDomainToModel(s)
	return h.db.WithContext(ctx).Create(&m).Error
}
