package images

import (
	imagesDomain "nfxidentity/modules/asset/domain/images"
	"nfxidentity/modules/asset/infrastructure/repository/images/check"
	"nfxidentity/modules/asset/infrastructure/repository/images/create"
	"nfxidentity/modules/asset/infrastructure/repository/images/delete"
	"nfxidentity/modules/asset/infrastructure/repository/images/get"
	"nfxidentity/modules/asset/infrastructure/repository/images/update"

	"gorm.io/gorm"
)

func NewRepo(db *gorm.DB) *imagesDomain.Repo {
	return &imagesDomain.Repo{
		Create: create.NewHandler(db),
		Get:    get.NewHandler(db),
		Check:  check.NewHandler(db),
		Update: update.NewHandler(db),
		Delete: delete.NewHandler(db),
	}
}
