package create

import (
	"context"
	authorizationCodeDomain "nfxid/modules/permission/domain/authorization_code"
	"nfxid/modules/permission/infrastructure/repository/mapper"
)

// New 创建新的 AuthorizationCode，实现 authorizationCodeDomain.Create 接口
func (h *Handler) New(ctx context.Context, ac *authorizationCodeDomain.AuthorizationCode) error {
	m := mapper.AuthorizationCodeDomainToModel(ac)
	return h.db.WithContext(ctx).Create(&m).Error
}
