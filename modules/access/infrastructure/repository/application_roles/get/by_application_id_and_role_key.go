package get

import (
	"context"
	"errors"

	accessErr "nfxidentity/errors/src/access"
	"nfxidentity/modules/access/domain/application_roles"
	"nfxidentity/modules/access/infrastructure/rdb/models"
	"nfxidentity/modules/access/infrastructure/repository/application_roles/mapper"

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
			return nil, accessErr.ErrApplicationRoleNotFound
		}
		return nil, err
	}
	return mapper.ApplicationRoleModelToDomain(&m), nil
}
