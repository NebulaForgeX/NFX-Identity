package delete

import (
	"context"
	"nfxid/modules/auth/domain/sessions"
	"nfxid/modules/auth/infrastructure/rdb/models"
)

// BySessionID 根据 SessionID 删除 Session，实现 sessions.Delete 接口
func (h *Handler) BySessionID(ctx context.Context, sessionID string) error {
	result := h.db.WithContext(ctx).
		Where("session_id = ?", sessionID).
		Delete(&models.Session{})

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return sessions.ErrSessionNotFound
	}
	return nil
}
