package get

import (
	"context"
	"errors"
	"nfxid/enums"
	"nfxid/modules/clients/domain/client_credentials"
	"nfxid/modules/clients/infrastructure/rdb/models"
	"nfxid/modules/clients/infrastructure/repository/client_credentials/mapper"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// ActiveByAppID 根据 AppID 获取活跃的 ClientCredential，实现 client_credentials.Get 接口
func (h *Handler) ActiveByAppID(ctx context.Context, appID uuid.UUID) (*client_credentials.ClientCredential, error) {
	var m models.ClientCredential
	if err := h.db.WithContext(ctx).
		Where("app_id = ? AND status = ?", appID, enums.ClientsCredentialStatusActive).
		First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, client_credentials.ErrClientCredentialNotFound
		}
		return nil, err
	}
	return mapper.ClientCredentialModelToDomain(&m), nil
}
