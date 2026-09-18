package files

import (
	filesDomain "nfxidentity/modules/asset/domain/files"
	"nfxidentity/modules/asset/infrastructure/repository/files/check"
	"nfxidentity/modules/asset/infrastructure/repository/files/create"
	"nfxidentity/modules/asset/infrastructure/repository/files/delete"
	"nfxidentity/modules/asset/infrastructure/repository/files/get"
	"nfxidentity/modules/asset/infrastructure/repository/files/update"

	"gorm.io/gorm"
)

func NewRepo(db *gorm.DB) *filesDomain.Repo {
	return &filesDomain.Repo{
		Create: create.NewHandler(db),
		Get:    get.NewHandler(db),
		Check:  check.NewHandler(db),
		Update: update.NewHandler(db),
		Delete: delete.NewHandler(db),
	}
}
