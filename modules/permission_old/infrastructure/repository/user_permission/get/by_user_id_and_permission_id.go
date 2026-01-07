package get

import (
	"context"
	"errors"
	userPermissionDomain "nfxid/modules/permission/domain/user_permission"
	userPermissionDomainErrors "nfxid/modules/permission/domain/user_permission/errors"
	"nfxid/modules/permission/infrastructure/rdb/models"
	"nfxid/modules/permission/infrastructure/repository/mapper"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// ByUserIDAndPermissionID 根据 UserID 和 PermissionID 获取 UserPermission，实现 userPermissionDomain.Get 接口
func (h *Handler) ByUserIDAndPermissionID(ctx context.Context, userID, permissionID uuid.UUID) (*userPermissionDomain.UserPermission, error) {
	var m models.UserPermission
	if err := h.db.WithContext(ctx).
		Where("user_id = ? AND permission_id = ?", userID, permissionID).
		Where("deleted_at IS NULL").
		First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, userPermissionDomainErrors.ErrUserPermissionNotFound
		}
		return nil, err
	}
	return mapper.UserPermissionModelToDomain(&m), nil
}
