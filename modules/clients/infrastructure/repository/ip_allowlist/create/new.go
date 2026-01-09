package create

import (
	"context"
	"nfxid/modules/clients/domain/ip_allowlist"
	"nfxid/modules/clients/infrastructure/repository/ip_allowlist/mapper"
)

// New 创建新的 IPAllowlist，实现 ip_allowlist.Create 接口
func (h *Handler) New(ctx context.Context, ip *ip_allowlist.IPAllowlist) error {
	m := mapper.IPAllowlistDomainToModel(ip)
	return h.db.WithContext(ctx).Create(&m).Error
}
