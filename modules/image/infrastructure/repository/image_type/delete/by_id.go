package delete

import (
	"context"
	"nfxid/modules/image/infrastructure/rdb/models"

	"github.com/google/uuid"
)

// ByID 根据 ID 删除 ImageType，实现 imageTypeDomain.Delete 接口
func (h *Handler) ByID(ctx context.Context, id uuid.UUID) error {
	return h.db.WithContext(ctx).Delete(&models.ImageType{}, "id = ?", id).Error
}
