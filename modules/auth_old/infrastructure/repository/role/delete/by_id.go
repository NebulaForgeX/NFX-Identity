package delete

import (
	"context"
	roleDomainErrors "nfxid/modules/auth/domain/role/errors"
	"nfxid/modules/auth/infrastructure/rdb/models"

	"github.com/google/uuid"
)

// ByID 根据 ID 删除 Role，实现 role.Delete 接口
func (h *Handler) ByID(ctx context.Context, id uuid.UUID) error {
	result := h.db.WithContext(ctx).
		Where("id = ?", id).
		Delete(&models.Role{})

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return roleDomainErrors.ErrRoleNotFound
	}
	return nil
}
