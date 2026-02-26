package user_credentials

import (
	"nfxidentity/modules/auth/domain/user_credentials"
	"nfxidentity/modules/auth/infrastructure/repository/user_credentials/check"
	"nfxidentity/modules/auth/infrastructure/repository/user_credentials/create"
	"nfxidentity/modules/auth/infrastructure/repository/user_credentials/delete"
	"nfxidentity/modules/auth/infrastructure/repository/user_credentials/get"
	"nfxidentity/modules/auth/infrastructure/repository/user_credentials/update"

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
