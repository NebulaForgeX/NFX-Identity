package get

import (
	"context"
	"errors"
	"nfxid/modules/tenants/domain/domain_verifications"
	"nfxid/modules/tenants/infrastructure/rdb/models"
	"nfxid/modules/tenants/infrastructure/repository/domain_verifications/mapper"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// ByID 根据 ID 获取 DomainVerification，实现 domain_verifications.Get 接口
func (h *Handler) ByID(ctx context.Context, id uuid.UUID) (*domain_verifications.DomainVerification, error) {
	var m models.DomainVerification
	if err := h.db.WithContext(ctx).Where("id = ?", id).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domain_verifications.ErrDomainVerificationNotFound
		}
		return nil, err
	}
	return mapper.DomainVerificationModelToDomain(&m), nil
}
