package update

import (
	"context"
	"nfxid/modules/clients/domain/ip_allowlist"
	"nfxid/modules/clients/infrastructure/rdb/models"
	"nfxid/modules/clients/infrastructure/repository/ip_allowlist/mapper"
)

// Generic 通用更新 IPAllowlist，实现 ip_allowlist.Update 接口
func (h *Handler) Generic(ctx context.Context, ip *ip_allowlist.IPAllowlist) error {
	m := mapper.IPAllowlistDomainToModel(ip)
	updates := mapper.IPAllowlistModelToUpdates(m)
	return h.db.WithContext(ctx).
		Model(&models.IpAllowlist{}).
		Where("id = ?", ip.ID()).
		Updates(updates).Error
}
