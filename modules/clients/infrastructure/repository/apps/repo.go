package apps

import (
	"nfxidentity/modules/clients/domain/apps"
	"nfxidentity/modules/clients/infrastructure/repository/apps/check"
	"nfxidentity/modules/clients/infrastructure/repository/apps/create"
	"nfxidentity/modules/clients/infrastructure/repository/apps/delete"
	"nfxidentity/modules/clients/infrastructure/repository/apps/get"
	"nfxidentity/modules/clients/infrastructure/repository/apps/update"

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
