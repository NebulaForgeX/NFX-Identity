package create

import (
	"context"
	"nfxid/modules/clients/domain/api_keys"
	"nfxid/modules/clients/infrastructure/repository/api_keys/mapper"
)

// New 创建新的 APIKey，实现 api_keys.Create 接口
func (h *Handler) New(ctx context.Context, ak *api_keys.APIKey) error {
	m := mapper.APIKeyDomainToModel(ak)
	return h.db.WithContext(ctx).Create(&m).Error
}
