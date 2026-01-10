package update

import (
	"context"
	"nfxid/modules/clients/domain/ip_allowlist"
	"nfxid/modules/clients/infrastructure/rdb/models"
	"nfxid/modules/clients/infrastructure/repository/ip_allowlist/mapper"
	"time"

	"github.com/google/uuid"
)

// Status 更新 IPAllowlist 状态，实现 ip_allowlist.Update 接口
func (h *Handler) Status(ctx context.Context, id uuid.UUID, status ip_allowlist.AllowlistStatus) error {
	statusEnum := mapper.AllowlistStatusDomainToEnum(status)
	updates := map[string]any{
		models.IpAllowlistCols.Status:    statusEnum,
		models.IpAllowlistCols.UpdatedAt: time.Now().UTC(),
	}

	return h.db.WithContext(ctx).
		Model(&models.IpAllowlist{}).
		Where("id = ?", id).
		Updates(updates).Error
}
