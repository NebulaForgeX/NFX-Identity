package delete

import (
	"context"
	"nfxid/modules/image/infrastructure/rdb/models"

	"github.com/google/uuid"
)

// ByImageIDAndTag 根据 ImageID 和 Tag 删除 ImageTag，实现 image_tags.Delete 接口
func (h *Handler) ByImageIDAndTag(ctx context.Context, imageID uuid.UUID, tag string) error {
	return h.db.WithContext(ctx).
		Where("image_id = ? AND tag = ?", imageID, tag).
		Delete(&models.ImageTag{}).Error
}
