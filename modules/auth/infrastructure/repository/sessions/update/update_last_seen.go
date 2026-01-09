package update

import (
	"context"
	"time"
	"nfxid/modules/auth/infrastructure/rdb/models"
)

// UpdateLastSeen 更新最后访问时间，实现 sessions.Update 接口
func (h *Handler) UpdateLastSeen(ctx context.Context, sessionID string) error {
	now := time.Now().UTC()
	updates := map[string]any{
		models.SessionCols.LastSeenAt: now,
		models.SessionCols.UpdatedAt:  now,
	}

	return h.db.WithContext(ctx).
		Model(&models.Session{}).
		Where("session_id = ?", sessionID).
		Updates(updates).Error
}
