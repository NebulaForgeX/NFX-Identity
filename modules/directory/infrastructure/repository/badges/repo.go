package badges

import (
	"nfxidentity/modules/directory/domain/badges"
	"nfxidentity/modules/directory/infrastructure/repository/badges/check"
	"nfxidentity/modules/directory/infrastructure/repository/badges/create"
	"nfxidentity/modules/directory/infrastructure/repository/badges/delete"
	"nfxidentity/modules/directory/infrastructure/repository/badges/get"
	"nfxidentity/modules/directory/infrastructure/repository/badges/update"

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
