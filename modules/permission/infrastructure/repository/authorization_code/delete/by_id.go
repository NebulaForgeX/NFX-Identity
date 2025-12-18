package delete

import (
	"context"
	"nfxid/modules/permission/infrastructure/rdb/models"

	"github.com/google/uuid"
)

// ByID 根据 ID 删除 AuthorizationCode，实现 authorizationCodeDomain.Delete 接口
func (h *Handler) ByID(ctx context.Context, id uuid.UUID) error {
	return h.db.WithContext(ctx).Delete(&models.AuthorizationCode{}, "id = ?", id).Error
}
