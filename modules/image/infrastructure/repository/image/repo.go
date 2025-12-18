package image

import (
	imageDomain "nfxid/modules/image/domain/image"
	"nfxid/modules/image/infrastructure/repository/image/check"
	"nfxid/modules/image/infrastructure/repository/image/create"
	"nfxid/modules/image/infrastructure/repository/image/delete"
	"nfxid/modules/image/infrastructure/repository/image/get"
	"nfxid/modules/image/infrastructure/repository/image/update"

	"gorm.io/gorm"
)

// NewRepo 创建一个新的 Image repository
func NewRepo(db *gorm.DB) *imageDomain.Repo {
	return &imageDomain.Repo{
		Create: create.NewHandler(db),
		Get:    get.NewHandler(db),
		Check:  check.NewHandler(db),
		Update: update.NewHandler(db),
		Delete: delete.NewHandler(db),
	}
}
