package image_types

import (
	"nfxidentity/modules/image/domain/image_types"
	"nfxidentity/modules/image/infrastructure/repository/image_types/check"
	"nfxidentity/modules/image/infrastructure/repository/image_types/create"
	"nfxidentity/modules/image/infrastructure/repository/image_types/delete"
	"nfxidentity/modules/image/infrastructure/repository/image_types/get"
	"nfxidentity/modules/image/infrastructure/repository/image_types/update"

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
