package check

import (
	"context"
	"nfxid/modules/auth/infrastructure/rdb/models"
)

// BySessionID 根据 SessionID 检查 Session 是否存在，实现 sessions.Check 接口
func (h *Handler) BySessionID(ctx context.Context, sessionID string) (bool, error) {
	var count int64
	err := h.db.WithContext(ctx).
		Model(&models.Session{}).
		Where("session_id = ?", sessionID).
		Count(&count).Error
	return count > 0, err
}
