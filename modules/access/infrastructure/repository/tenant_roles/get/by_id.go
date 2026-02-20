package get

import (
	"context"
	"errors"

	"nfxid/modules/access/domain/tenant_roles"
	"nfxid/modules/access/infrastructure/rdb/models"
	"nfxid/modules/access/infrastructure/repository/tenant_roles/mapper"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// ByID 按 ID 获取
func (h *Handler) ByID(ctx context.Context, id uuid.UUID) (*tenant_roles.TenantRole, error) {
	var m models.TenantRole
	if err := h.db.WithContext(ctx).Where("id = ?", id).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, tenant_roles.ErrTenantRoleNotFound
		}
		return nil, err
	}
	return mapper.TenantRoleModelToDomain(&m), nil
}
