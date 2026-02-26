package update

import (
	"context"
	"nfxidentity/modules/tenants/domain/domain_verifications"
	"nfxidentity/modules/tenants/infrastructure/rdb/models"
	"nfxidentity/modules/tenants/infrastructure/repository/domain_verifications/mapper"
)

// Generic 通用更新 DomainVerification，实现 domain_verifications.Update 接口
func (h *Handler) Generic(ctx context.Context, dv *domain_verifications.DomainVerification) error {
	m := mapper.DomainVerificationDomainToModel(dv)
	updates := mapper.DomainVerificationModelToUpdates(m)
	return h.db.WithContext(ctx).
		Model(&models.DomainVerification{}).
		Where("id = ?", dv.ID()).
		Updates(updates).Error
}
