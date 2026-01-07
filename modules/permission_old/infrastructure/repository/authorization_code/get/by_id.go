package get

import (
	"context"
	"errors"
	authorizationCodeDomain "nfxid/modules/permission/domain/authorization_code"
	authorizationCodeDomainErrors "nfxid/modules/permission/domain/authorization_code/errors"
	"nfxid/modules/permission/infrastructure/rdb/models"
	"nfxid/modules/permission/infrastructure/repository/mapper"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// ByID 根据 ID 获取 AuthorizationCode，实现 authorizationCodeDomain.Get 接口
func (h *Handler) ByID(ctx context.Context, id uuid.UUID) (*authorizationCodeDomain.AuthorizationCode, error) {
	var m models.AuthorizationCode
	if err := h.db.WithContext(ctx).Where("id = ?", id).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, authorizationCodeDomainErrors.ErrAuthorizationCodeNotFound
		}
		return nil, err
	}
	return mapper.AuthorizationCodeModelToDomain(&m), nil
}
