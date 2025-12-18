package create

import (
	"context"
	userPermissionDomain "nfxid/modules/permission/domain/user_permission"
	"nfxid/modules/permission/infrastructure/repository/mapper"
)

// New 创建新的 UserPermission，实现 userPermissionDomain.Create 接口
func (h *Handler) New(ctx context.Context, up *userPermissionDomain.UserPermission) error {
	m := mapper.UserPermissionDomainToModel(up)
	return h.db.WithContext(ctx).Create(&m).Error
}
