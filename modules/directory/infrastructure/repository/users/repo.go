package users

import (
	"nfxid/modules/directory/domain/users"
	"nfxid/modules/directory/infrastructure/repository/users/check"
	"nfxid/modules/directory/infrastructure/repository/users/create"
	"nfxid/modules/directory/infrastructure/repository/users/delete"
	"nfxid/modules/directory/infrastructure/repository/users/get"
	"nfxid/modules/directory/infrastructure/repository/users/update"

	"gorm.io/gorm"
)

// NewRepo 创建一个新的 User repository
func NewRepo(db *gorm.DB) *users.Repo {
	return &users.Repo{
		Create: create.NewHandler(db),
		Get:    get.NewHandler(db),
		Check:  check.NewHandler(db),
		Update: update.NewHandler(db),
		Delete: delete.NewHandler(db),
	}
}
