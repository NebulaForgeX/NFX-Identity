package get

import (
	"context"
	"errors"

	accessErr "nfxidentity/errors/src/access"
	"nfxidentity/modules/access/domain/tenant_roles"
	"nfxidentity/modules/access/infrastructure/rdb/models"
	"nfxidentity/modules/access/infrastructure/repository/tenant_roles/mapper"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// ByID 按 ID 获取
func (h *Handler) ByID(ctx context.Context, id uuid.UUID) (*tenant_roles.TenantRole, error) {
	var m models.TenantRole
	if err := h.db.WithContext(ctx).Where("id = ?", id).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, accessErr.ErrTenantRoleNotFound
		}
		return nil, err
	}
	return mapper.TenantRoleModelToDomain(&m), nil
}
