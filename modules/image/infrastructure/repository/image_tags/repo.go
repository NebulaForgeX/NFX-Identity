package image_tags

import (
	"nfxidentity/modules/image/domain/image_tags"
	"nfxidentity/modules/image/infrastructure/repository/image_tags/check"
	"nfxidentity/modules/image/infrastructure/repository/image_tags/create"
	"nfxidentity/modules/image/infrastructure/repository/image_tags/delete"
	"nfxidentity/modules/image/infrastructure/repository/image_tags/get"
	"nfxidentity/modules/image/infrastructure/repository/image_tags/update"

	"gorm.io/gorm"
)

// NewRepo 创建一个新的 ImageTag repository
func NewRepo(db *gorm.DB) *image_tags.Repo {
	return &image_tags.Repo{
		Create: create.NewHandler(db),
		Get:    get.NewHandler(db),
		Check:  check.NewHandler(db),
		Update: update.NewHandler(db),
		Delete: delete.NewHandler(db),
	}
}
