package profile_badge

import (
	profileBadge "nfxid/modules/auth/domain/profile_badge"
	"nfxid/modules/auth/infrastructure/repository/profile_badge/check"
	"nfxid/modules/auth/infrastructure/repository/profile_badge/create"
	"nfxid/modules/auth/infrastructure/repository/profile_badge/delete"
	"nfxid/modules/auth/infrastructure/repository/profile_badge/get"
	"nfxid/modules/auth/infrastructure/repository/profile_badge/update"

	"gorm.io/gorm"
)

// NewRepo 创建一个新的 ProfileBadge repository
func NewRepo(db *gorm.DB) *profileBadge.Repo {
	return &profileBadge.Repo{
		Create: create.NewHandler(db),
		Get:    get.NewHandler(db),
		Check:  check.NewHandler(db),
		Update: update.NewHandler(db),
		Delete: delete.NewHandler(db),
	}
}
