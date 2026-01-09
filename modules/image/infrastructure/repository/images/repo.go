package images

import (
	"nfxid/modules/image/domain/images"
	"nfxid/modules/image/infrastructure/repository/images/check"
	"nfxid/modules/image/infrastructure/repository/images/create"
	"nfxid/modules/image/infrastructure/repository/images/delete"
	"nfxid/modules/image/infrastructure/repository/images/get"
	"nfxid/modules/image/infrastructure/repository/images/update"

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
