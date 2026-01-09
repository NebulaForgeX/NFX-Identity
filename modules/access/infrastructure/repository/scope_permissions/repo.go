package scope_permissions

import (
	"nfxid/modules/access/domain/scope_permissions"
	"nfxid/modules/access/infrastructure/repository/scope_permissions/check"
	"nfxid/modules/access/infrastructure/repository/scope_permissions/create"
	"nfxid/modules/access/infrastructure/repository/scope_permissions/delete"
	"nfxid/modules/access/infrastructure/repository/scope_permissions/get"
	"nfxid/modules/access/infrastructure/repository/scope_permissions/update"

	"gorm.io/gorm"
)

// NewRepo 创建一个新的 ScopePermission repository
func NewRepo(db *gorm.DB) *scope_permissions.Repo {
	return &scope_permissions.Repo{
		Create: create.NewHandler(db),
		Get:    get.NewHandler(db),
		Check:  check.NewHandler(db),
		Update: update.NewHandler(db),
		Delete: delete.NewHandler(db),
	}
}
