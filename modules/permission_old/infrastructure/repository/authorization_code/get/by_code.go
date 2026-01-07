package get

import (
	"context"
	"errors"
	authorizationCodeDomain "nfxid/modules/permission/domain/authorization_code"
	authorizationCodeDomainErrors "nfxid/modules/permission/domain/authorization_code/errors"
	"nfxid/modules/permission/infrastructure/rdb/models"
	"nfxid/modules/permission/infrastructure/repository/mapper"

	"gorm.io/gorm"
)

// ByCode 根据 Code 获取 AuthorizationCode，实现 authorizationCodeDomain.Get 接口
func (h *Handler) ByCode(ctx context.Context, code string) (*authorizationCodeDomain.AuthorizationCode, error) {
	var m models.AuthorizationCode
	if err := h.db.WithContext(ctx).Where("code = ?", code).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, authorizationCodeDomainErrors.ErrAuthorizationCodeNotFound
		}
		return nil, err
	}
	return mapper.AuthorizationCodeModelToDomain(&m), nil
}
