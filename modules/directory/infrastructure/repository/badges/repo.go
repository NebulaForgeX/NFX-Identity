package badges

import (
	"nfxid/modules/directory/domain/badges"
	"nfxid/modules/directory/infrastructure/repository/badges/check"
	"nfxid/modules/directory/infrastructure/repository/badges/create"
	"nfxid/modules/directory/infrastructure/repository/badges/delete"
	"nfxid/modules/directory/infrastructure/repository/badges/get"
	"nfxid/modules/directory/infrastructure/repository/badges/update"

	"gorm.io/gorm"
)

// NewRepo 创建一个新的 Badge repository
func NewRepo(db *gorm.DB) *badges.Repo {
	return &badges.Repo{
		Create: create.NewHandler(db),
		Get:    get.NewHandler(db),
		Check:  check.NewHandler(db),
		Update: update.NewHandler(db),
		Delete: delete.NewHandler(db),
	}
}
