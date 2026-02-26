package user_profiles

import (
	"nfxidentity/modules/directory/domain/user_profiles"
	"nfxidentity/modules/directory/infrastructure/repository/user_profiles/check"
	"nfxidentity/modules/directory/infrastructure/repository/user_profiles/create"
	"nfxidentity/modules/directory/infrastructure/repository/user_profiles/delete"
	"nfxidentity/modules/directory/infrastructure/repository/user_profiles/get"
	"nfxidentity/modules/directory/infrastructure/repository/user_profiles/update"

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
