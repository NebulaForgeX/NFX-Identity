package image_tags

import (
	"nfxid/modules/image/domain/image_tags"
	"nfxid/modules/image/infrastructure/repository/image_tags/check"
	"nfxid/modules/image/infrastructure/repository/image_tags/create"
	"nfxid/modules/image/infrastructure/repository/image_tags/delete"
	"nfxid/modules/image/infrastructure/repository/image_tags/get"
	"nfxid/modules/image/infrastructure/repository/image_tags/update"

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
