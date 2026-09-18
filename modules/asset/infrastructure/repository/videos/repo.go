package videos

import (
	videosDomain "nfxidentity/modules/asset/domain/videos"
	"nfxidentity/modules/asset/infrastructure/repository/videos/check"
	"nfxidentity/modules/asset/infrastructure/repository/videos/create"
	"nfxidentity/modules/asset/infrastructure/repository/videos/delete"
	"nfxidentity/modules/asset/infrastructure/repository/videos/get"
	"nfxidentity/modules/asset/infrastructure/repository/videos/update"

	"gorm.io/gorm"
)

func NewRepo(db *gorm.DB) *videosDomain.Repo {
	return &videosDomain.Repo{
		Create: create.NewHandler(db),
		Get:    get.NewHandler(db),
		Check:  check.NewHandler(db),
		Update: update.NewHandler(db),
		Delete: delete.NewHandler(db),
	}
}
