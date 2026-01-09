package check

import (
	"context"
	"nfxid/modules/tenants/infrastructure/rdb/models"

	"github.com/google/uuid"
)

// ByTenantIDAndDomain 根据 TenantID 和 Domain 检查 DomainVerification 是否存在，实现 domain_verifications.Check 接口
func (h *Handler) ByTenantIDAndDomain(ctx context.Context, tenantID uuid.UUID, domain string) (bool, error) {
	var count int64
	err := h.db.WithContext(ctx).
		Model(&models.DomainVerification{}).
		Where("tenant_id = ? AND domain = ?", tenantID, domain).
		Count(&count).Error
	return count > 0, err
}
