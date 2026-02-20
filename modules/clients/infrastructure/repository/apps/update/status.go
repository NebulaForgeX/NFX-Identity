package update

import (
	"context"
	"nfxid/modules/clients/domain/apps"
	"nfxid/modules/clients/infrastructure/rdb/models"
	"nfxid/modules/clients/infrastructure/repository/apps/mapper"
	"time"

	"github.com/google/uuid"
)

// Status 更新 App 状态，实现 apps.Update 接口
func (h *Handler) Status(ctx context.Context, id uuid.UUID, status apps.AppStatus) error {
	statusEnum := mapper.AppStatusDomainToEnum(status)
	updates := map[string]any{
		models.ApplicationCols.Status:    statusEnum,
		models.ApplicationCols.UpdatedAt: time.Now().UTC(),
	}

	return h.db.WithContext(ctx).
		Model(&models.Application{}).
		Where("id = ?", id).
		Updates(updates).Error
}
