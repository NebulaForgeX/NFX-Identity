package profile

import (
	profileDomain "nfxidentity/modules/auth/domain/profile"
	"nfxidentity/modules/auth/infrastructure/repository/profile/check"
	"nfxidentity/modules/auth/infrastructure/repository/profile/create"
	"nfxidentity/modules/auth/infrastructure/repository/profile/delete"
	"nfxidentity/modules/auth/infrastructure/repository/profile/get"
	"nfxidentity/modules/auth/infrastructure/repository/profile/update"

	"gorm.io/gorm"
)

func NewRepo(db *gorm.DB) *profileDomain.Repo {
	return &profileDomain.Repo{
		Create: create.NewHandler(db),
		Get:    get.NewHandler(db),
		Check:  check.NewHandler(db),
		Update: update.NewHandler(db),
		Delete: delete.NewHandler(db),
	}
}
