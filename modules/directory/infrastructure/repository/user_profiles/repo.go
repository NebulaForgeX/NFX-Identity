package user_profiles

import (
	"nfxid/modules/directory/domain/user_profiles"
	"nfxid/modules/directory/infrastructure/repository/user_profiles/check"
	"nfxid/modules/directory/infrastructure/repository/user_profiles/create"
	"nfxid/modules/directory/infrastructure/repository/user_profiles/delete"
	"nfxid/modules/directory/infrastructure/repository/user_profiles/get"
	"nfxid/modules/directory/infrastructure/repository/user_profiles/update"

	"gorm.io/gorm"
)

// NewRepo 创建一个新的 UserProfile repository
func NewRepo(db *gorm.DB) *user_profiles.Repo {
	return &user_profiles.Repo{
		Create: create.NewHandler(db),
		Get:    get.NewHandler(db),
		Check:  check.NewHandler(db),
		Update: update.NewHandler(db),
		Delete: delete.NewHandler(db),
	}
}
