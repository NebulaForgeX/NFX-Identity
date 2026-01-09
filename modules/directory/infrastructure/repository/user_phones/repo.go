package user_phones

import (
	"nfxid/modules/directory/domain/user_phones"
	"nfxid/modules/directory/infrastructure/repository/user_phones/check"
	"nfxid/modules/directory/infrastructure/repository/user_phones/create"
	"nfxid/modules/directory/infrastructure/repository/user_phones/delete"
	"nfxid/modules/directory/infrastructure/repository/user_phones/get"
	"nfxid/modules/directory/infrastructure/repository/user_phones/update"

	"gorm.io/gorm"
)

// NewRepo 创建一个新的 UserPhone repository
func NewRepo(db *gorm.DB) *user_phones.Repo {
	return &user_phones.Repo{
		Create: create.NewHandler(db),
		Get:    get.NewHandler(db),
		Check:  check.NewHandler(db),
		Update: update.NewHandler(db),
		Delete: delete.NewHandler(db),
	}
}
