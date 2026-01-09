package role_permissions

import (
	"nfxid/modules/access/domain/role_permissions"
	"nfxid/modules/access/infrastructure/repository/role_permissions/check"
	"nfxid/modules/access/infrastructure/repository/role_permissions/create"
	"nfxid/modules/access/infrastructure/repository/role_permissions/delete"
	"nfxid/modules/access/infrastructure/repository/role_permissions/get"
	"nfxid/modules/access/infrastructure/repository/role_permissions/update"

	"gorm.io/gorm"
)

// NewRepo 创建一个新的 RolePermission repository
func NewRepo(db *gorm.DB) *role_permissions.Repo {
	return &role_permissions.Repo{
		Create: create.NewHandler(db),
		Get:    get.NewHandler(db),
		Check:  check.NewHandler(db),
		Update: update.NewHandler(db),
		Delete: delete.NewHandler(db),
	}
}
