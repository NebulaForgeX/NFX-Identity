package user_badges

import (
	"nfxidentity/modules/directory/domain/user_badges"
	"nfxidentity/modules/directory/infrastructure/repository/user_badges/check"
	"nfxidentity/modules/directory/infrastructure/repository/user_badges/create"
	"nfxidentity/modules/directory/infrastructure/repository/user_badges/delete"
	"nfxidentity/modules/directory/infrastructure/repository/user_badges/get"
	"nfxidentity/modules/directory/infrastructure/repository/user_badges/update"

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
