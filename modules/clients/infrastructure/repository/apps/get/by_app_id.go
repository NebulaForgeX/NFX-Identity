package get

import (
	"context"
	"errors"

	clientsErr "nfxidentity/errors/src/clients"
	"nfxidentity/modules/clients/domain/apps"
	"nfxidentity/modules/clients/infrastructure/rdb/models"
	"nfxidentity/modules/clients/infrastructure/repository/apps/mapper"

	"gorm.io/gorm"
)

// ByAppID 根据 AppID 获取 App，实现 apps.Get 接口
func (h *Handler) ByAppID(ctx context.Context, appID string) (*apps.App, error) {
	var m models.Application
	if err := h.db.WithContext(ctx).Where("application_id = ?", appID).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, clientsErr.ErrAppNotFound
		}
		return nil, err
	}
	return mapper.AppModelToDomain(&m), nil
}
