package single

import (
	"context"
	"errors"
	userDomainErrors "nfxid/modules/auth/domain/user/errors"
	userDomainViews "nfxid/modules/auth/domain/user/views"
	"nfxid/modules/auth/infrastructure/query/mapper"
	"nfxid/modules/auth/infrastructure/rdb/models"
	"nfxid/modules/auth/infrastructure/rdb/views"

	"gorm.io/gorm"
)

// ByEmail 根据邮箱获取 User，实现 userDomain.Single 接口
func (h *Handler) ByEmail(ctx context.Context, email string) (*userDomainViews.UserView, error) {
	var v views.UserWithRoleView
	if err := h.db.WithContext(ctx).Where("email = ?", email).First(&v).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, userDomainErrors.ErrUserViewNotFound
		}
		return nil, err
	}
	var u models.User
	if err := h.db.WithContext(ctx).Where("email = ?", email).First(&u).Error; err != nil {
		return nil, err
	}
	view := mapper.UserViewToDomain(&v, &u)
	return &view, nil
}

