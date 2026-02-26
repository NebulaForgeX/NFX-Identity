package user_preferences

import (
	"nfxidentity/modules/directory/domain/user_preferences"
	"nfxidentity/modules/directory/infrastructure/repository/user_preferences/check"
	"nfxidentity/modules/directory/infrastructure/repository/user_preferences/create"
	"nfxidentity/modules/directory/infrastructure/repository/user_preferences/delete"
	"nfxidentity/modules/directory/infrastructure/repository/user_preferences/get"
	"nfxidentity/modules/directory/infrastructure/repository/user_preferences/update"

	"gorm.io/gorm"
)

// NewRepo 创建一个新的 UserPreference repository
func NewRepo(db *gorm.DB) *user_preferences.Repo {
	return &user_preferences.Repo{
		Create: create.NewHandler(db),
		Get:    get.NewHandler(db),
		Check:  check.NewHandler(db),
		Update: update.NewHandler(db),
		Delete: delete.NewHandler(db),
	}
}
