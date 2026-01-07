package single

import (
	"context"
	"errors"
	userRoleDomainErrors "nfxid/modules/auth/domain/user_role/errors"
	userRoleDomainViews "nfxid/modules/auth/domain/user_role/views"
	"nfxid/modules/auth/infrastructure/query/mapper"
	"nfxid/modules/auth/infrastructure/rdb/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// ByID 根据 ID 获取 UserRole，实现 userRoleDomain.Single 接口
func (h *Handler) ByID(ctx context.Context, id uuid.UUID) (*userRoleDomainViews.UserRoleView, error) {
	var m models.UserRole
	if err := h.db.WithContext(ctx).Where("id = ?", id).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, userRoleDomainErrors.ErrUserRoleViewNotFound
		}
		return nil, err
	}
	view := mapper.UserRoleModelToDomain(&m)
	return &view, nil
}
