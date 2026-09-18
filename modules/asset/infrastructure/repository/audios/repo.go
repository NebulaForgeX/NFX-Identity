package audios

import (
	audiosDomain "nfxidentity/modules/asset/domain/audios"
	"nfxidentity/modules/asset/infrastructure/repository/audios/check"
	"nfxidentity/modules/asset/infrastructure/repository/audios/create"
	"nfxidentity/modules/asset/infrastructure/repository/audios/delete"
	"nfxidentity/modules/asset/infrastructure/repository/audios/get"
	"nfxidentity/modules/asset/infrastructure/repository/audios/update"

	"gorm.io/gorm"
)

func NewRepo(db *gorm.DB) *audiosDomain.Repo {
	return &audiosDomain.Repo{
		Create: create.NewHandler(db),
		Get:    get.NewHandler(db),
		Check:  check.NewHandler(db),
		Update: update.NewHandler(db),
		Delete: delete.NewHandler(db),
	}
}
