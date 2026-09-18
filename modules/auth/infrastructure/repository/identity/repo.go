package identity

import (
	identityDomain "nfxidentity/modules/auth/domain/identity"
	"nfxidentity/modules/auth/infrastructure/repository/identity/check"
	"nfxidentity/modules/auth/infrastructure/repository/identity/create"
	"nfxidentity/modules/auth/infrastructure/repository/identity/delete"
	"nfxidentity/modules/auth/infrastructure/repository/identity/get"
	"nfxidentity/modules/auth/infrastructure/repository/identity/update"

	"gorm.io/gorm"
)

func NewRepo(db *gorm.DB) *identityDomain.Repo {
	return &identityDomain.Repo{
		Create: create.NewHandler(db),
		Get:    get.NewHandler(db),
		Check:  check.NewHandler(db),
		Update: update.NewHandler(db),
		Delete: delete.NewHandler(db),
	}
}
