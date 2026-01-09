package get

import (
	"context"
	"errors"
	"nfxid/modules/clients/domain/client_credentials"
	"nfxid/modules/clients/infrastructure/rdb/models"
	"nfxid/modules/clients/infrastructure/repository/client_credentials/mapper"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// ByID 根据 ID 获取 ClientCredential，实现 client_credentials.Get 接口
func (h *Handler) ByID(ctx context.Context, id uuid.UUID) (*client_credentials.ClientCredential, error) {
	var m models.ClientCredential
	if err := h.db.WithContext(ctx).Where("id = ?", id).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, client_credentials.ErrClientCredentialNotFound
		}
		return nil, err
	}
	return mapper.ClientCredentialModelToDomain(&m), nil
}
