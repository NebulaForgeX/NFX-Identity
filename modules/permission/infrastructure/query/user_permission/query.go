package user_permission

import (
	userPermissionDomain "nfxid/modules/permission/domain/user_permission"
	"nfxid/modules/permission/infrastructure/query/user_permission/list"
	"nfxid/modules/permission/infrastructure/query/user_permission/single"

	"gorm.io/gorm"
)

// NewHandler 创建新的 UserPermission Query Handler
func NewHandler(db *gorm.DB) *userPermissionDomain.Query {
	return &userPermissionDomain.Query{
		Single: single.NewHandler(db),
		List:   list.NewHandler(db),
	}
}
