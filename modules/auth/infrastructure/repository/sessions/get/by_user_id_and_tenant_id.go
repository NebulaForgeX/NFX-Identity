package get

import (
	"context"
	"nfxid/modules/auth/domain/sessions"
	"nfxid/modules/auth/infrastructure/rdb/models"
	"nfxid/modules/auth/infrastructure/repository/sessions/mapper"

	"github.com/google/uuid"
)

// ByUserIDAndTenantID 根据 UserID 和 TenantID 获取 Session 列表，实现 sessions.Get 接口
func (h *Handler) ByUserIDAndTenantID(ctx context.Context, userID, tenantID uuid.UUID) ([]*sessions.Session, error) {
	var ms []models.Session
	if err := h.db.WithContext(ctx).
		Where("user_id = ? AND tenant_id = ?", userID, tenantID).
		Find(&ms).Error; err != nil {
		return nil, err
	}

	result := make([]*sessions.Session, len(ms))
	for i, m := range ms {
		result[i] = mapper.SessionModelToDomain(&m)
	}
	return result, nil
}
