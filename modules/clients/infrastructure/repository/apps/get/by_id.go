package get

import (
	"context"
	"errors"
	"nfxid/modules/clients/domain/apps"
	"nfxid/modules/clients/infrastructure/rdb/models"
	"nfxid/modules/clients/infrastructure/repository/apps/mapper"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// ByID 根据 ID 获取 App，实现 apps.Get 接口
func (h *Handler) ByID(ctx context.Context, id uuid.UUID) (*apps.App, error) {
	var m models.Application
	if err := h.db.WithContext(ctx).Where("id = ?", id).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, apps.ErrAppNotFound
		}
		return nil, err
	}
	return mapper.AppModelToDomain(&m), nil
}
