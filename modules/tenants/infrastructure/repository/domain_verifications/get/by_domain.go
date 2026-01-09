package get

import (
	"context"
	"errors"
	"nfxid/modules/tenants/domain/domain_verifications"
	"nfxid/modules/tenants/infrastructure/rdb/models"
	"nfxid/modules/tenants/infrastructure/repository/domain_verifications/mapper"

	"gorm.io/gorm"
)

// ByDomain 根据 Domain 获取 DomainVerification，实现 domain_verifications.Get 接口
func (h *Handler) ByDomain(ctx context.Context, domain string) (*domain_verifications.DomainVerification, error) {
	var m models.DomainVerification
	if err := h.db.WithContext(ctx).Where("domain = ?", domain).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domain_verifications.ErrDomainVerificationNotFound
		}
		return nil, err
	}
	return mapper.DomainVerificationModelToDomain(&m), nil
}
