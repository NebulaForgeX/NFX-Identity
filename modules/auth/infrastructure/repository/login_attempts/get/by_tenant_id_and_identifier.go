package get

import (
	"context"
	"nfxid/modules/auth/domain/login_attempts"
	"nfxid/modules/auth/infrastructure/rdb/models"
	"nfxid/modules/auth/infrastructure/repository/login_attempts/mapper"

	"github.com/google/uuid"
)

// ByTenantIDAndIdentifier 根据 TenantID 和 Identifier 获取 LoginAttempt 列表，实现 login_attempts.Get 接口
func (h *Handler) ByTenantIDAndIdentifier(ctx context.Context, tenantID uuid.UUID, identifier string) ([]*login_attempts.LoginAttempt, error) {
	var ms []models.LoginAttempt
	if err := h.db.WithContext(ctx).
		Where("tenant_id = ? AND identifier = ?", tenantID, identifier).
		Find(&ms).Error; err != nil {
		return nil, err
	}

	result := make([]*login_attempts.LoginAttempt, len(ms))
	for i, m := range ms {
		result[i] = mapper.LoginAttemptModelToDomain(&m)
	}
	return result, nil
}
