package create

import (
	"context"
	"nfxidentity/modules/auth/domain/sessions"
	"nfxidentity/modules/auth/infrastructure/repository/sessions/mapper"
)

// New 创建新的 Session，实现 sessions.Create 接口
func (h *Handler) New(ctx context.Context, s *sessions.Session) error {
	m := mapper.SessionDomainToModel(s)
	return h.db.WithContext(ctx).Create(&m).Error
}
