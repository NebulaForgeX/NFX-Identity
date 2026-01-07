package user

import (
	"nfxid/modules/auth/domain/user"
	"nfxid/modules/auth/infrastructure/repository/user/check"
	"nfxid/modules/auth/infrastructure/repository/user/create"
	"nfxid/modules/auth/infrastructure/repository/user/delete"
	"nfxid/modules/auth/infrastructure/repository/user/get"
	"nfxid/modules/auth/infrastructure/repository/user/update"

	"gorm.io/gorm"
)

// NewRepo 创建一个新的 User repository
func NewRepo(db *gorm.DB) *user.Repo {
	return &user.Repo{
		Create: create.NewHandler(db),
		Get:    get.NewHandler(db),
		Check:  check.NewHandler(db),
		Update: update.NewHandler(db),
		Delete: delete.NewHandler(db),
	}
}
