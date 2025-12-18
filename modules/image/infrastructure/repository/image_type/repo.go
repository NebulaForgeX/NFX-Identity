package image_type

import (
	imageTypeDomain "nfxid/modules/image/domain/image_type"
	"nfxid/modules/image/infrastructure/repository/image_type/check"
	"nfxid/modules/image/infrastructure/repository/image_type/create"
	"nfxid/modules/image/infrastructure/repository/image_type/delete"
	"nfxid/modules/image/infrastructure/repository/image_type/get"
	"nfxid/modules/image/infrastructure/repository/image_type/update"

	"gorm.io/gorm"
)

// NewRepo 创建一个新的 ImageType repository
func NewRepo(db *gorm.DB) *imageTypeDomain.Repo {
	return &imageTypeDomain.Repo{
		Create: create.NewHandler(db),
		Get:    get.NewHandler(db),
		Check:  check.NewHandler(db),
		Update: update.NewHandler(db),
		Delete: delete.NewHandler(db),
	}
}
