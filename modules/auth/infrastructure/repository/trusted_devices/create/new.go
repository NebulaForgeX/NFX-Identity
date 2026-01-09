package create

import (
	"context"
	"nfxid/modules/auth/domain/trusted_devices"
	"nfxid/modules/auth/infrastructure/repository/trusted_devices/mapper"
)

// New 创建新的 TrustedDevice，实现 trusted_devices.Create 接口
func (h *Handler) New(ctx context.Context, td *trusted_devices.TrustedDevice) error {
	m := mapper.TrustedDeviceDomainToModel(td)
	return h.db.WithContext(ctx).Create(&m).Error
}
