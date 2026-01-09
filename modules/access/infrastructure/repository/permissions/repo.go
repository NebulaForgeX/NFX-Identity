package permissions

import (
	"nfxid/modules/access/domain/permissions"
	"nfxid/modules/access/infrastructure/repository/permissions/check"
	"nfxid/modules/access/infrastructure/repository/permissions/create"
	"nfxid/modules/access/infrastructure/repository/permissions/delete"
	"nfxid/modules/access/infrastructure/repository/permissions/get"
	"nfxid/modules/access/infrastructure/repository/permissions/update"

	"gorm.io/gorm"
)

// NewRepo 创建一个新的 Permission repository
func NewRepo(db *gorm.DB) *permissions.Repo {
	return &permissions.Repo{
		Create: create.NewHandler(db),
		Get:    get.NewHandler(db),
		Check:  check.NewHandler(db),
		Update: update.NewHandler(db),
		Delete: delete.NewHandler(db),
	}
}
