package delete

import (
	"context"
	"errors"
	"nfxid/modules/clients/domain/ip_allowlist"
	"nfxid/modules/clients/infrastructure/rdb/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// ByID 根据 ID 删除 IPAllowlist，实现 ip_allowlist.Delete 接口
func (h *Handler) ByID(ctx context.Context, id uuid.UUID) error {
	result := h.db.WithContext(ctx).
		Where("id = ?", id).
		Delete(&models.IpAllowlist{})

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return ip_allowlist.ErrIPAllowlistNotFound
	}
	return nil
}
