package get

import (
	"context"
	"errors"

	accessErr "nfxidentity/errors/src/access"
	"nfxidentity/modules/access/domain/tenant_role_assignments"
	"nfxidentity/modules/access/infrastructure/rdb/models"
	"nfxidentity/modules/access/infrastructure/repository/tenant_role_assignments/mapper"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (h *Handler) ByID(ctx context.Context, id uuid.UUID) (*tenant_role_assignments.TenantRoleAssignment, error) {
	var m models.TenantRoleAssignment
	if err := h.db.WithContext(ctx).Where("id = ?", id).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, accessErr.ErrTenantRoleAssignmentNotFound
		}
		return nil, err
	}
	return mapper.TenantRoleAssignmentModelToDomain(&m), nil
}
