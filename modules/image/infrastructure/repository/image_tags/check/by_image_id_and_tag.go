package check

import (
	"context"
	"nfxid/modules/image/infrastructure/rdb/models"

	"github.com/google/uuid"
)

// ByImageIDAndTag 根据 ImageID 和 Tag 检查 ImageTag 是否存在，实现 image_tags.Check 接口
func (h *Handler) ByImageIDAndTag(ctx context.Context, imageID uuid.UUID, tag string) (bool, error) {
	var count int64
	err := h.db.WithContext(ctx).
		Model(&models.ImageTag{}).
		Where("image_id = ? AND tag = ?", imageID, tag).
		Count(&count).Error
	return count > 0, err
}
