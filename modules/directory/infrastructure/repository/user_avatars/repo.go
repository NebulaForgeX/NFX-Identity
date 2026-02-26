package user_avatars

import (
	"nfxidentity/modules/directory/domain/user_avatars"
	"nfxidentity/modules/directory/infrastructure/repository/user_avatars/check"
	"nfxidentity/modules/directory/infrastructure/repository/user_avatars/create"
	"nfxidentity/modules/directory/infrastructure/repository/user_avatars/delete"
	"nfxidentity/modules/directory/infrastructure/repository/user_avatars/get"
	"nfxidentity/modules/directory/infrastructure/repository/user_avatars/update"

	"gorm.io/gorm"
)

// NewRepo 创建一个新的 UserAvatar repository
func NewRepo(db *gorm.DB) *user_avatars.Repo {
	return &user_avatars.Repo{
		Create: create.NewHandler(db),
		Get:    get.NewHandler(db),
		Check:  check.NewHandler(db),
		Update: update.NewHandler(db),
		Delete: delete.NewHandler(db),
	}
}
