package phone

import (
	phoneDomain "nfxidentity/modules/auth/domain/phone"
	"nfxidentity/modules/auth/infrastructure/repository/phone/check"
	"nfxidentity/modules/auth/infrastructure/repository/phone/create"
	"nfxidentity/modules/auth/infrastructure/repository/phone/delete"
	"nfxidentity/modules/auth/infrastructure/repository/phone/get"
	"nfxidentity/modules/auth/infrastructure/repository/phone/update"

	"gorm.io/gorm"
)

func NewRepo(db *gorm.DB) *phoneDomain.Repo {
	return &phoneDomain.Repo{
		Create: create.NewHandler(db),
		Get:    get.NewHandler(db),
		Check:  check.NewHandler(db),
		Update: update.NewHandler(db),
		Delete: delete.NewHandler(db),
	}
}
