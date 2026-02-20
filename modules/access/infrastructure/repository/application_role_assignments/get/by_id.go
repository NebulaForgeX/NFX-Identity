package get

import (
	"context"
	"errors"

	"nfxid/modules/access/domain/application_role_assignments"
	"nfxid/modules/access/infrastructure/rdb/models"
	"nfxid/modules/access/infrastructure/repository/application_role_assignments/mapper"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (h *Handler) ByID(ctx context.Context, id uuid.UUID) (*application_role_assignments.ApplicationRoleAssignment, error) {
	var m models.ApplicationRoleAssignment
	if err := h.db.WithContext(ctx).Where("id = ?", id).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, application_role_assignments.ErrApplicationRoleAssignmentNotFound
		}
		return nil, err
	}
	return mapper.ApplicationRoleAssignmentModelToDomain(&m), nil
}
