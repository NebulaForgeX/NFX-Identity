package permission

import (
	permissionDomain "nfxid/modules/permission/domain/permission"
	"nfxid/modules/permission/infrastructure/repository/permission/check"
	"nfxid/modules/permission/infrastructure/repository/permission/create"
	"nfxid/modules/permission/infrastructure/repository/permission/delete"
	"nfxid/modules/permission/infrastructure/repository/permission/get"
	"nfxid/modules/permission/infrastructure/repository/permission/update"

	"gorm.io/gorm"
)

// NewRepo 创建一个新的 Permission repository
func NewRepo(db *gorm.DB) *permissionDomain.Repo {
	return &permissionDomain.Repo{
		Create: create.NewHandler(db),
		Get:    get.NewHandler(db),
		Check:  check.NewHandler(db),
		Update: update.NewHandler(db),
		Delete: delete.NewHandler(db),
	}
}
