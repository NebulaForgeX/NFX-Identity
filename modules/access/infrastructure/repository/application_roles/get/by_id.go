package get

import (
	"context"
	"errors"

	"nfxid/modules/access/domain/application_roles"
	"nfxid/modules/access/infrastructure/rdb/models"
	"nfxid/modules/access/infrastructure/repository/application_roles/mapper"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// ByID 按 ID 获取
func (h *Handler) ByID(ctx context.Context, id uuid.UUID) (*application_roles.ApplicationRole, error) {
	var m models.ApplicationRole
	if err := h.db.WithContext(ctx).Where("id = ?", id).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, application_roles.ErrApplicationRoleNotFound
		}
		return nil, err
	}
	return mapper.ApplicationRoleModelToDomain(&m), nil
}
