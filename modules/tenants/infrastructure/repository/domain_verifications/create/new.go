package create

import (
	"context"
	"nfxidentity/modules/tenants/domain/domain_verifications"
	"nfxidentity/modules/tenants/infrastructure/repository/domain_verifications/mapper"
)

// New 创建新的 DomainVerification，实现 domain_verifications.Create 接口
func (h *Handler) New(ctx context.Context, dv *domain_verifications.DomainVerification) error {
	m := mapper.DomainVerificationDomainToModel(dv)
	return h.db.WithContext(ctx).Create(&m).Error
}
