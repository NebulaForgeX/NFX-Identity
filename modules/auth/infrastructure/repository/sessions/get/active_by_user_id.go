package get

import (
	"context"
	"time"
	"nfxid/modules/auth/domain/sessions"
	"nfxid/modules/auth/infrastructure/rdb/models"
	"nfxid/modules/auth/infrastructure/repository/sessions/mapper"

	"github.com/google/uuid"
)

// ActiveByUserID 根据 UserID 和 TenantID 获取活跃的 Session 列表，实现 sessions.Get 接口
func (h *Handler) ActiveByUserID(ctx context.Context, userID, tenantID uuid.UUID) ([]*sessions.Session, error) {
	now := time.Now().UTC()
	var ms []models.Session
	if err := h.db.WithContext(ctx).
		Where("user_id = ? AND tenant_id = ? AND expires_at > ? AND revoked_at IS NULL", userID, tenantID, now).
		Find(&ms).Error; err != nil {
		return nil, err
	}

	result := make([]*sessions.Session, len(ms))
	for i, m := range ms {
		result[i] = mapper.SessionModelToDomain(&m)
	}
	return result, nil
}
