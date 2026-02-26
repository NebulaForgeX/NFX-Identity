package users

import (
	"nfxidentity/modules/directory/domain/users"
	"nfxidentity/modules/directory/infrastructure/repository/users/check"
	"nfxidentity/modules/directory/infrastructure/repository/users/create"
	"nfxidentity/modules/directory/infrastructure/repository/users/delete"
	"nfxidentity/modules/directory/infrastructure/repository/users/get"
	"nfxidentity/modules/directory/infrastructure/repository/users/update"

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
