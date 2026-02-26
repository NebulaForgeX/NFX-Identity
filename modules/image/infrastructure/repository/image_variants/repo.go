package image_variants

import (
	"nfxidentity/modules/image/domain/image_variants"
	"nfxidentity/modules/image/infrastructure/repository/image_variants/check"
	"nfxidentity/modules/image/infrastructure/repository/image_variants/create"
	"nfxidentity/modules/image/infrastructure/repository/image_variants/delete"
	"nfxidentity/modules/image/infrastructure/repository/image_variants/get"
	"nfxidentity/modules/image/infrastructure/repository/image_variants/update"

	"gorm.io/gorm"
)

// NewRepo 创建一个新的 ImageVariant repository
func NewRepo(db *gorm.DB) *image_variants.Repo {
	return &image_variants.Repo{
		Create: create.NewHandler(db),
		Get:    get.NewHandler(db),
		Check:  check.NewHandler(db),
		Update: update.NewHandler(db),
		Delete: delete.NewHandler(db),
	}
}
