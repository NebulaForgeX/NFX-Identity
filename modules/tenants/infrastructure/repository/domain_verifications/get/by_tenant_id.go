package get

import (
	"context"
	"nfxidentity/modules/tenants/domain/domain_verifications"
	"nfxidentity/modules/tenants/infrastructure/rdb/models"
	"nfxidentity/modules/tenants/infrastructure/repository/domain_verifications/mapper"

	"github.com/google/uuid"
)

// ByTenantID 根据 TenantID 获取 DomainVerification 列表，实现 domain_verifications.Get 接口
func (h *Handler) ByTenantID(ctx context.Context, tenantID uuid.UUID) ([]*domain_verifications.DomainVerification, error) {
	var ms []models.DomainVerification
	if err := h.db.WithContext(ctx).
		Where("tenant_id = ?", tenantID).
		Find(&ms).Error; err != nil {
		return nil, err
	}

	result := make([]*domain_verifications.DomainVerification, len(ms))
	for i, m := range ms {
		result[i] = mapper.DomainVerificationModelToDomain(&m)
	}
	return result, nil
}
