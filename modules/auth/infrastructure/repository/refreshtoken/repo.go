package refreshtoken

import (
	refreshtokenDomain "nfxidentity/modules/auth/domain/refreshtoken"
	"nfxidentity/modules/auth/infrastructure/repository/refreshtoken/check"
	"nfxidentity/modules/auth/infrastructure/repository/refreshtoken/create"
	"nfxidentity/modules/auth/infrastructure/repository/refreshtoken/delete"
	"nfxidentity/modules/auth/infrastructure/repository/refreshtoken/get"
	"nfxidentity/modules/auth/infrastructure/repository/refreshtoken/update"

	"gorm.io/gorm"
)

func NewRepo(db *gorm.DB) *refreshtokenDomain.Repo {
	return &refreshtokenDomain.Repo{
		Create: create.NewHandler(db),
		Get:    get.NewHandler(db),
		Check:  check.NewHandler(db),
		Update: update.NewHandler(db),
		Delete: delete.NewHandler(db),
	}
}
