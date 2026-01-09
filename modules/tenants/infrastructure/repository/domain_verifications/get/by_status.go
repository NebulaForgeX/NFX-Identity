package get

import (
	"context"
	"nfxid/modules/tenants/domain/domain_verifications"
	"nfxid/modules/tenants/infrastructure/rdb/models"
	"nfxid/modules/tenants/infrastructure/repository/domain_verifications/mapper"
)

// ByStatus 根据 Status 获取 DomainVerification 列表，实现 domain_verifications.Get 接口
func (h *Handler) ByStatus(ctx context.Context, status domain_verifications.VerificationStatus) ([]*domain_verifications.DomainVerification, error) {
	statusEnum := mapper.VerificationStatusDomainToEnum(status)
	var ms []models.DomainVerification
	if err := h.db.WithContext(ctx).
		Where("status = ?", statusEnum).
		Find(&ms).Error; err != nil {
		return nil, err
	}

	result := make([]*domain_verifications.DomainVerification, len(ms))
	for i, m := range ms {
		result[i] = mapper.DomainVerificationModelToDomain(&m)
	}
	return result, nil
}
