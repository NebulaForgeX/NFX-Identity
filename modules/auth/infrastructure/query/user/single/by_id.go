package single

import (
	"context"
	"errors"
	userDomainErrors "nfxid/modules/auth/domain/user/errors"
	userDomainViews "nfxid/modules/auth/domain/user/views"
	"nfxid/modules/auth/infrastructure/query/mapper"
	"nfxid/modules/auth/infrastructure/rdb/models"
	"nfxid/modules/auth/infrastructure/rdb/views"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// ByID 根据 ID 获取 User，实现 userDomain.Single 接口
func (h *Handler) ByID(ctx context.Context, id uuid.UUID) (*userDomainViews.UserView, error) {
	var v views.UserWithRoleView
	if err := h.db.WithContext(ctx).Where("user_id = ?", id).First(&v).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, userDomainErrors.ErrUserViewNotFound
		}
		return nil, err
	}
	var u models.User
	if err := h.db.WithContext(ctx).Where("id = ?", id).First(&u).Error; err != nil {
		return nil, err
	}
	view := mapper.UserViewToDomain(&v, &u)
	return &view, nil
}

