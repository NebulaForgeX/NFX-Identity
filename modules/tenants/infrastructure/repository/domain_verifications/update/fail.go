package update

import (
	"context"
	"errors"
	"nfxid/enums"
	"nfxid/modules/tenants/domain/domain_verifications"
	"nfxid/modules/tenants/infrastructure/rdb/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Fail 标记验证失败，实现 domain_verifications.Update 接口
func (h *Handler) Fail(ctx context.Context, id uuid.UUID) error {
	// 先检查 DomainVerification 是否存在
	var m models.DomainVerification
	if err := h.db.WithContext(ctx).
		Where("id = ?", id).
		First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return domain_verifications.ErrDomainVerificationNotFound
		}
		return err
	}

	status := enums.TenantsVerificationStatusFailed
	updates := map[string]any{
		models.DomainVerificationCols.Status: status,
	}

	return h.db.WithContext(ctx).
		Model(&models.DomainVerification{}).
		Where("id = ?", id).
		Updates(updates).Error
}
