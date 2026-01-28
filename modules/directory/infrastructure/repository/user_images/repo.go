package user_images

import (
	"nfxid/modules/directory/domain/user_images"
	"nfxid/modules/directory/infrastructure/repository/user_images/check"
	"nfxid/modules/directory/infrastructure/repository/user_images/create"
	"nfxid/modules/directory/infrastructure/repository/user_images/delete"
	"nfxid/modules/directory/infrastructure/repository/user_images/get"
	"nfxid/modules/directory/infrastructure/repository/user_images/update"

	"gorm.io/gorm"
)

// NewRepo 创建一个新的 UserImage repository
func NewRepo(db *gorm.DB) *user_images.Repo {
	return &user_images.Repo{
		Create: create.NewHandler(db),
		Get:    get.NewHandler(db),
		Check:  check.NewHandler(db),
		Update: update.NewHandler(db),
		Delete: delete.NewHandler(db),
	}
}
