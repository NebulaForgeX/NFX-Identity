package get

import (
	"context"
	"errors"
	"nfxid/modules/auth/domain/role"
	roleDomainErrors "nfxid/modules/auth/domain/role/errors"
	"nfxid/modules/auth/infrastructure/rdb/models"
	"nfxid/modules/auth/infrastructure/repository/mapper"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// ByID 根据 ID 获取 Role，实现 role.Get 接口
func (h *Handler) ByID(ctx context.Context, id uuid.UUID) (*role.Role, error) {
	var m models.Role
	if err := h.db.WithContext(ctx).Where("id = ?", id).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, roleDomainErrors.ErrRoleNotFound
		}
		return nil, err
	}
	return mapper.RoleModelToDomain(&m), nil
}
