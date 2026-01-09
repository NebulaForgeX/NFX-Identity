package image_variants

import (
	"nfxid/modules/image/domain/image_variants"
	"nfxid/modules/image/infrastructure/repository/image_variants/check"
	"nfxid/modules/image/infrastructure/repository/image_variants/create"
	"nfxid/modules/image/infrastructure/repository/image_variants/delete"
	"nfxid/modules/image/infrastructure/repository/image_variants/get"
	"nfxid/modules/image/infrastructure/repository/image_variants/update"

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
