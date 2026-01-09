package apps

import (
	"nfxid/modules/clients/domain/apps"
	"nfxid/modules/clients/infrastructure/repository/apps/check"
	"nfxid/modules/clients/infrastructure/repository/apps/create"
	"nfxid/modules/clients/infrastructure/repository/apps/delete"
	"nfxid/modules/clients/infrastructure/repository/apps/get"
	"nfxid/modules/clients/infrastructure/repository/apps/update"

	"gorm.io/gorm"
)

// NewRepo 创建一个新的 App repository
func NewRepo(db *gorm.DB) *apps.Repo {
	return &apps.Repo{
		Create: create.NewHandler(db),
		Get:    get.NewHandler(db),
		Check:  check.NewHandler(db),
		Update: update.NewHandler(db),
		Delete: delete.NewHandler(db),
	}
}
