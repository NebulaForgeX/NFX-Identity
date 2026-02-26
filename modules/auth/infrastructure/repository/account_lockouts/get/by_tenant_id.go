package get

import (
	"context"
	"nfxidentity/modules/auth/domain/account_lockouts"
	"nfxidentity/modules/auth/infrastructure/rdb/models"
	"nfxidentity/modules/auth/infrastructure/repository/account_lockouts/mapper"

	"github.com/google/uuid"
)

// ByTenantID 根据 TenantID 获取 AccountLockout 列表，实现 account_lockouts.Get 接口
func (h *Handler) ByTenantID(ctx context.Context, tenantID uuid.UUID) ([]*account_lockouts.AccountLockout, error) {
	var ms []models.AccountLockout
	if err := h.db.WithContext(ctx).
		Where("tenant_id = ?", tenantID).
		Find(&ms).Error; err != nil {
		return nil, err
	}

	result := make([]*account_lockouts.AccountLockout, len(ms))
	for i, m := range ms {
		result[i] = mapper.AccountLockoutModelToDomain(&m)
	}
	return result, nil
}
