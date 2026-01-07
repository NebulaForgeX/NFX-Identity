package delete

import (
	"context"
	"nfxid/modules/image/infrastructure/rdb/models"
)

// ByStoragePath 根据 StoragePath 删除 Image，实现 imageDomain.Delete 接口
func (h *Handler) ByStoragePath(ctx context.Context, storagePath string) error {
	return h.db.WithContext(ctx).Delete(&models.Image{}, "storage_path = ?", storagePath).Error
}
