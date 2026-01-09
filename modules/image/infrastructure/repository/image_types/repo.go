package image_types

import (
	"nfxid/modules/image/domain/image_types"
	"nfxid/modules/image/infrastructure/repository/image_types/check"
	"nfxid/modules/image/infrastructure/repository/image_types/create"
	"nfxid/modules/image/infrastructure/repository/image_types/delete"
	"nfxid/modules/image/infrastructure/repository/image_types/get"
	"nfxid/modules/image/infrastructure/repository/image_types/update"

	"gorm.io/gorm"
)

// NewRepo 创建一个新的 ImageType repository
func NewRepo(db *gorm.DB) *image_types.Repo {
	return &image_types.Repo{
		Create: create.NewHandler(db),
		Get:    get.NewHandler(db),
		Check:  check.NewHandler(db),
		Update: update.NewHandler(db),
		Delete: delete.NewHandler(db),
	}
}
