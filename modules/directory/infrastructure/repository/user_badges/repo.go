package user_badges

import (
	"nfxid/modules/directory/domain/user_badges"
	"nfxid/modules/directory/infrastructure/repository/user_badges/check"
	"nfxid/modules/directory/infrastructure/repository/user_badges/create"
	"nfxid/modules/directory/infrastructure/repository/user_badges/delete"
	"nfxid/modules/directory/infrastructure/repository/user_badges/get"
	"nfxid/modules/directory/infrastructure/repository/user_badges/update"

	"gorm.io/gorm"
)

// NewRepo 创建一个新的 UserBadge repository
func NewRepo(db *gorm.DB) *user_badges.Repo {
	return &user_badges.Repo{
		Create: create.NewHandler(db),
		Get:    get.NewHandler(db),
		Check:  check.NewHandler(db),
		Update: update.NewHandler(db),
		Delete: delete.NewHandler(db),
	}
}
