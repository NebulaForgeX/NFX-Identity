package update

import (
	"context"
	authorizationCodeDomain "nfxid/modules/permission/domain/authorization_code"
	"nfxid/modules/permission/infrastructure/rdb/models"
	"nfxid/modules/permission/infrastructure/repository/mapper"
)

// Generic 通用更新 AuthorizationCode，实现 authorizationCodeDomain.Update 接口
func (h *Handler) Generic(ctx context.Context, ac *authorizationCodeDomain.AuthorizationCode) error {
	m := mapper.AuthorizationCodeDomainToModel(ac)
	updates := mapper.AuthorizationCodeModelsToUpdates(m)
	return h.db.WithContext(ctx).Model(&models.AuthorizationCode{}).Where("id = ?", ac.ID()).Updates(updates).Error
}
