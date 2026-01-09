package get

import (
	"context"
	"errors"
	"nfxid/modules/auth/domain/sessions"
	"nfxid/modules/auth/infrastructure/rdb/models"
	"nfxid/modules/auth/infrastructure/repository/sessions/mapper"

	"gorm.io/gorm"
)

// BySessionID 根据 SessionID 获取 Session，实现 sessions.Get 接口
func (h *Handler) BySessionID(ctx context.Context, sessionID string) (*sessions.Session, error) {
	var m models.Session
	if err := h.db.WithContext(ctx).Where("session_id = ?", sessionID).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, sessions.ErrSessionNotFound
		}
		return nil, err
	}
	return mapper.SessionModelToDomain(&m), nil
}
