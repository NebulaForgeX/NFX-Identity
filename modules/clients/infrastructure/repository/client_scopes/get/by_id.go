package get

import (
	"context"
	"errors"

	clientsErr "nfxidentity/errors/src/clients"
	"nfxidentity/modules/clients/domain/client_scopes"
	"nfxidentity/modules/clients/infrastructure/rdb/models"
	"nfxidentity/modules/clients/infrastructure/repository/client_scopes/mapper"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// ByID 根据 ID 获取 ClientScope，实现 client_scopes.Get 接口
func (h *Handler) ByID(ctx context.Context, id uuid.UUID) (*client_scopes.ClientScope, error) {
	var m models.ClientScope
	if err := h.db.WithContext(ctx).Where("id = ?", id).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, clientsErr.ErrClientScopeNotFound
		}
		return nil, err
	}
	return mapper.ClientScopeModelToDomain(&m), nil
}
