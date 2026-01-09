package update

import (
	"context"
	"nfxid/modules/auth/domain/sessions"
	"nfxid/modules/auth/infrastructure/rdb/models"
	"nfxid/modules/auth/infrastructure/repository/sessions/mapper"
)

// Generic 通用更新 Session，实现 sessions.Update 接口
func (h *Handler) Generic(ctx context.Context, s *sessions.Session) error {
	m := mapper.SessionDomainToModel(s)
	updates := mapper.SessionModelToUpdates(m)
	return h.db.WithContext(ctx).
		Model(&models.Session{}).
		Where("id = ?", s.ID()).
		Updates(updates).Error
}
