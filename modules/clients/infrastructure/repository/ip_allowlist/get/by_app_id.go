package get

import (
	"context"
	"nfxid/modules/clients/domain/ip_allowlist"
	"nfxid/modules/clients/infrastructure/rdb/models"
	"nfxid/modules/clients/infrastructure/repository/ip_allowlist/mapper"

	"github.com/google/uuid"
)

// ByAppID 根据 AppID 获取 IPAllowlist 列表，实现 ip_allowlist.Get 接口
func (h *Handler) ByAppID(ctx context.Context, appID uuid.UUID) ([]*ip_allowlist.IPAllowlist, error) {
	var ms []models.IpAllowlist
	if err := h.db.WithContext(ctx).
		Where("app_id = ?", appID).
		Find(&ms).Error; err != nil {
		return nil, err
	}

	result := make([]*ip_allowlist.IPAllowlist, len(ms))
	for i, m := range ms {
		result[i] = mapper.IPAllowlistModelToDomain(&m)
	}
	return result, nil
}
