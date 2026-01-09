package update

import (
	"context"
	"nfxid/modules/access/domain/scopes"
	"nfxid/modules/access/infrastructure/rdb/models"
	"nfxid/modules/access/infrastructure/repository/scopes/mapper"
)

// Generic 通用更新 Scope，实现 scopes.Update 接口
func (h *Handler) Generic(ctx context.Context, s *scopes.Scope) error {
	m := mapper.ScopeDomainToModel(s)
	updates := mapper.ScopeModelToUpdates(m)
	return h.db.WithContext(ctx).
		Model(&models.Scope{}).
		Where("scope = ?", s.ScopeKey()).
		Updates(updates).Error
}
