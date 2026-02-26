package user_images

import (
	"nfxidentity/modules/directory/domain/user_images"
	"nfxidentity/modules/directory/infrastructure/repository/user_images/check"
	"nfxidentity/modules/directory/infrastructure/repository/user_images/create"
	"nfxidentity/modules/directory/infrastructure/repository/user_images/delete"
	"nfxidentity/modules/directory/infrastructure/repository/user_images/get"
	"nfxidentity/modules/directory/infrastructure/repository/user_images/update"

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
