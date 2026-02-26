package get

import (
	"context"
	"errors"
	clientsErr "nfxidentity/errors/src/clients"
	"nfxidentity/modules/clients/domain/client_credentials"
	"nfxidentity/modules/clients/infrastructure/rdb/models"
	"nfxidentity/modules/clients/infrastructure/repository/client_credentials/mapper"

	"gorm.io/gorm"
)

// ByClientID 根据 ClientID 获取 ClientCredential，实现 client_credentials.Get 接口
func (h *Handler) ByClientID(ctx context.Context, clientID string) (*client_credentials.ClientCredential, error) {
	var m models.ClientCredential
	if err := h.db.WithContext(ctx).Where("client_id = ?", clientID).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, clientsErr.ErrClientCredentialNotFound
		}
		return nil, err
	}
	return mapper.ClientCredentialModelToDomain(&m), nil
}
