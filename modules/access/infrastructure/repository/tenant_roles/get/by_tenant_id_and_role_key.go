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

// ByTenantIDAndRoleKey 按租户ID与角色键获取
func (h *Handler) ByTenantIDAndRoleKey(ctx context.Context, tenantID uuid.UUID, roleKey string) (*tenant_roles.TenantRole, error) {
	var m models.TenantRole
	if err := h.db.WithContext(ctx).
		Where("tenant_id = ? AND role_key = ?", tenantID, roleKey).
		First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, tenant_roles.ErrTenantRoleNotFound
		}
		return nil, err
	}
	return mapper.TenantRoleModelToDomain(&m), nil
}
