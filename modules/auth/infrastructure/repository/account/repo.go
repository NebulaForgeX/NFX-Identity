package account

import (
	accountDomain "nfxidentity/modules/auth/domain/account"
	"nfxidentity/modules/auth/infrastructure/repository/account/check"
	"nfxidentity/modules/auth/infrastructure/repository/account/create"
	"nfxidentity/modules/auth/infrastructure/repository/account/delete"
	"nfxidentity/modules/auth/infrastructure/repository/account/get"
	"nfxidentity/modules/auth/infrastructure/repository/account/update"

	"gorm.io/gorm"
)

func NewRepo(db *gorm.DB) *accountDomain.Repo {
	return &accountDomain.Repo{
		Create: create.NewHandler(db),
		Get:    get.NewHandler(db),
		Check:  check.NewHandler(db),
		Update: update.NewHandler(db),
		Delete: delete.NewHandler(db),
	}
}
