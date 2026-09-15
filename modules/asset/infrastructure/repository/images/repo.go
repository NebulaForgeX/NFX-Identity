package images

import (
	"nfxidentity/modules/asset/domain/images"
	"nfxidentity/modules/asset/infrastructure/repository/images/create"
	"nfxidentity/modules/asset/infrastructure/repository/images/get"
	"nfxidentity/modules/asset/infrastructure/repository/images/update"

	"gorm.io/gorm"
)

func NewRepo(db *gorm.DB) *images.Repo {
	return &images.Repo{
		Create: create.NewHandler(db),
		Get:    get.NewHandler(db),
		Update: update.NewHandler(db),
	}
}
