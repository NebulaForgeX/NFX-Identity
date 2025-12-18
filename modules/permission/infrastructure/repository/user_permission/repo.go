package user_permission

import (
	userPermissionDomain "nfxid/modules/permission/domain/user_permission"
	"nfxid/modules/permission/infrastructure/repository/user_permission/check"
	"nfxid/modules/permission/infrastructure/repository/user_permission/create"
	"nfxid/modules/permission/infrastructure/repository/user_permission/delete"
	"nfxid/modules/permission/infrastructure/repository/user_permission/get"

	"gorm.io/gorm"
)

// NewRepo 创建一个新的 UserPermission repository
func NewRepo(db *gorm.DB) *userPermissionDomain.Repo {
	return &userPermissionDomain.Repo{
		Create: create.NewHandler(db),
		Get:    get.NewHandler(db),
		Check:  check.NewHandler(db),
		Delete: delete.NewHandler(db),
	}
}
