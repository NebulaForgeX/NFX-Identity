package role

import (
	"context"
	"errors"
	roleDomainErrors "nfxid/modules/auth/domain/role/errors"
	roleDomainViews "nfxid/modules/auth/domain/role/views"
	"nfxid/modules/auth/infrastructure/query/mapper"
	"nfxid/modules/auth/infrastructure/rdb/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// ByID 根据 ID 获取 Role，实现 role.Query 接口
func (h *Handler) ByID(ctx context.Context, id uuid.UUID) (roleDomainViews.RoleView, error) {
	var m models.Role
	if err := h.db.WithContext(ctx).Where("id = ?", id).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return roleDomainViews.RoleView{}, roleDomainErrors.ErrRoleViewNotFound
		}
		return roleDomainViews.RoleView{}, err
	}
	return mapper.RoleModelToDomain(&m), nil
}
