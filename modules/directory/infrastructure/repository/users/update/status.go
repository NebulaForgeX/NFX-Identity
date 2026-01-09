package update

import (
	"context"
	"time"
	"nfxid/modules/directory/domain/users"
	"nfxid/modules/directory/infrastructure/rdb/models"
	"nfxid/modules/directory/infrastructure/repository/users/mapper"

	"github.com/google/uuid"
)

// Status 更新状态，实现 users.Update 接口
func (h *Handler) Status(ctx context.Context, id uuid.UUID, status users.UserStatus) error {
	statusEnum := mapper.UserStatusDomainToEnum(status)
	updates := map[string]any{
		models.UserCols.Status:    statusEnum,
		models.UserCols.UpdatedAt: time.Now().UTC(),
	}

	return h.db.WithContext(ctx).
		Model(&models.User{}).
		Where("id = ?", id).
		Updates(updates).Error
}
