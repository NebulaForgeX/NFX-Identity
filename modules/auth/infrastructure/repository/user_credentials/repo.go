package user_credentials

import (
	"nfxid/modules/auth/domain/user_credentials"
	"nfxid/modules/auth/infrastructure/repository/user_credentials/check"
	"nfxid/modules/auth/infrastructure/repository/user_credentials/create"
	"nfxid/modules/auth/infrastructure/repository/user_credentials/delete"
	"nfxid/modules/auth/infrastructure/repository/user_credentials/get"
	"nfxid/modules/auth/infrastructure/repository/user_credentials/update"

	"gorm.io/gorm"
)

// NewRepo 创建一个新的 UserCredential repository
func NewRepo(db *gorm.DB) *user_credentials.Repo {
	return &user_credentials.Repo{
		Create: create.NewHandler(db),
		Get:    get.NewHandler(db),
		Check:  check.NewHandler(db),
		Update: update.NewHandler(db),
		Delete: delete.NewHandler(db),
	}
}
