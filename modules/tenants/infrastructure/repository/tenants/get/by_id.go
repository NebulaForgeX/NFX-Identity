package get

import (
	"context"
	"errors"
	"nfxid/modules/tenants/domain/tenants"
	"nfxid/modules/tenants/infrastructure/rdb/models"
	"nfxid/modules/tenants/infrastructure/repository/tenants/mapper"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// ByID 根据 ID 获取 Tenant，实现 tenants.Get 接口
func (h *Handler) ByID(ctx context.Context, id uuid.UUID) (*tenants.Tenant, error) {
	var m models.Tenant
	if err := h.db.WithContext(ctx).Where("id = ?", id).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, tenants.ErrTenantNotFound
		}
		return nil, err
	}
	return mapper.TenantModelToDomain(&m), nil
}
