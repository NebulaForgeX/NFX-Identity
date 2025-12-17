package profile

import (
	"nfxid/modules/auth/domain/profile"
	"nfxid/modules/auth/infrastructure/repository/profile/check"
	"nfxid/modules/auth/infrastructure/repository/profile/create"
	"nfxid/modules/auth/infrastructure/repository/profile/delete"
	"nfxid/modules/auth/infrastructure/repository/profile/get"
	"nfxid/modules/auth/infrastructure/repository/profile/update"

	"gorm.io/gorm"
)

// NewRepo 创建一个新的 Profile repository
func NewRepo(db *gorm.DB) *profile.Repo {
	return &profile.Repo{
		Create: create.NewHandler(db),
		Get:    get.NewHandler(db),
		Check:  check.NewHandler(db),
		Update: update.NewHandler(db),
		Delete: delete.NewHandler(db),
	}
}
