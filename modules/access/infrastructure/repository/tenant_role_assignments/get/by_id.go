package get

import (
	"context"
	"errors"

	"nfxid/modules/access/domain/tenant_role_assignments"
	"nfxid/modules/access/infrastructure/rdb/models"
	"nfxid/modules/access/infrastructure/repository/tenant_role_assignments/mapper"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (h *Handler) ByID(ctx context.Context, id uuid.UUID) (*tenant_role_assignments.TenantRoleAssignment, error) {
	var m models.TenantRoleAssignment
	if err := h.db.WithContext(ctx).Where("id = ?", id).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, tenant_role_assignments.ErrTenantRoleAssignmentNotFound
		}
		return nil, err
	}
	return mapper.TenantRoleAssignmentModelToDomain(&m), nil
}
