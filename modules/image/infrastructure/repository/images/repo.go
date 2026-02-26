package images

import (
	"nfxidentity/modules/image/domain/images"
	"nfxidentity/modules/image/infrastructure/repository/images/check"
	"nfxidentity/modules/image/infrastructure/repository/images/create"
	"nfxidentity/modules/image/infrastructure/repository/images/delete"
	"nfxidentity/modules/image/infrastructure/repository/images/get"
	"nfxidentity/modules/image/infrastructure/repository/images/update"

	"gorm.io/gorm"
)

// NewRepo 创建一个新的 Image repository
func NewRepo(db *gorm.DB) *images.Repo {
	return &images.Repo{
		Create: create.NewHandler(db),
		Get:    get.NewHandler(db),
		Check:  check.NewHandler(db),
		Update: update.NewHandler(db),
		Delete: delete.NewHandler(db),
	}
}
