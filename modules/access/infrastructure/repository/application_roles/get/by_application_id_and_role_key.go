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

// ByApplicationIDAndRoleKey 按应用ID与角色键获取
func (h *Handler) ByApplicationIDAndRoleKey(ctx context.Context, applicationID uuid.UUID, roleKey string) (*application_roles.ApplicationRole, error) {
	var m models.ApplicationRole
	if err := h.db.WithContext(ctx).
		Where("application_id = ? AND role_key = ?", applicationID, roleKey).
		First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, application_roles.ErrApplicationRoleNotFound
		}
		return nil, err
	}
	return mapper.ApplicationRoleModelToDomain(&m), nil
}
