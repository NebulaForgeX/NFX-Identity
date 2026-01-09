package update

import (
	"context"
	"errors"
	"time"
	"nfxid/modules/auth/domain/sessions"
	"nfxid/modules/auth/infrastructure/rdb/models"
	"nfxid/modules/auth/infrastructure/repository/sessions/mapper"

	"gorm.io/gorm"
)

// Revoke 撤销 Session，实现 sessions.Update 接口
func (h *Handler) Revoke(ctx context.Context, sessionID string, reason sessions.SessionRevokeReason, revokedBy string) error {
	// 先检查 Session 是否存在
	var m models.Session
	if err := h.db.WithContext(ctx).
		Where("session_id = ?", sessionID).
		First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return sessions.ErrSessionNotFound
		}
		return err
	}

	// 检查是否已经撤销
	if m.RevokedAt != nil {
		return sessions.ErrSessionAlreadyRevoked
	}

	now := time.Now().UTC()
	revokeReason := mapper.SessionRevokeReasonDomainToEnum(reason)
	updates := map[string]any{
		models.SessionCols.RevokedAt:    &now,
		models.SessionCols.RevokeReason: &revokeReason,
		models.SessionCols.RevokedBy:    &revokedBy,
		models.SessionCols.UpdatedAt:    now,
	}

	return h.db.WithContext(ctx).
		Model(&models.Session{}).
		Where("session_id = ?", sessionID).
		Updates(updates).Error
}
