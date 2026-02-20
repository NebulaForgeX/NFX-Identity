package get

import (
	"context"

	"nfxid/modules/access/domain/application_roles"
	"nfxid/modules/access/infrastructure/rdb/models"
	"nfxid/modules/access/infrastructure/repository/application_roles/mapper"

	"github.com/google/uuid"
)

// ByApplicationID 按应用ID列表
func (h *Handler) ByApplicationID(ctx context.Context, applicationID uuid.UUID) ([]*application_roles.ApplicationRole, error) {
	var list []*models.ApplicationRole
	if err := h.db.WithContext(ctx).
		Where("application_id = ?", applicationID).
		Order("created_at ASC").
		Find(&list).Error; err != nil {
		return nil, err
	}
	out := make([]*application_roles.ApplicationRole, len(list))
	for i := range list {
		out[i] = mapper.ApplicationRoleModelToDomain(list[i])
	}
	return out, nil
}
