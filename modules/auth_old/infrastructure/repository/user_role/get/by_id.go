package get

import (
	"context"
	"errors"
	"nfxid/modules/auth/domain/user_role"
	userRoleErrors "nfxid/modules/auth/domain/user_role/errors"
	"nfxid/modules/auth/infrastructure/rdb/models"
	"nfxid/modules/auth/infrastructure/repository/mapper"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// ByID 根据 ID 获取 UserRole，实现 user_role.Get 接口
func (h *Handler) ByID(ctx context.Context, id uuid.UUID) (*user_role.UserRole, error) {
	var m models.UserRole
	if err := h.db.WithContext(ctx).Where("id = ?", id).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, userRoleErrors.ErrUserRoleNotFound
		}
		return nil, err
	}
	return mapper.UserRoleModelToDomain(&m), nil
}
