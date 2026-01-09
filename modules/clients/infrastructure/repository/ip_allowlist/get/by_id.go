package get

import (
	"context"
	"errors"
	"nfxid/modules/clients/domain/ip_allowlist"
	"nfxid/modules/clients/infrastructure/rdb/models"
	"nfxid/modules/clients/infrastructure/repository/ip_allowlist/mapper"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// ByID 根据 ID 获取 IPAllowlist，实现 ip_allowlist.Get 接口
func (h *Handler) ByID(ctx context.Context, id uuid.UUID) (*ip_allowlist.IPAllowlist, error) {
	var m models.IpAllowlist
	if err := h.db.WithContext(ctx).Where("id = ?", id).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ip_allowlist.ErrIPAllowlistNotFound
		}
		return nil, err
	}
	return mapper.IPAllowlistModelToDomain(&m), nil
}
