package user_preferences

import (
	"nfxid/modules/directory/domain/user_preferences"
	"nfxid/modules/directory/infrastructure/repository/user_preferences/check"
	"nfxid/modules/directory/infrastructure/repository/user_preferences/create"
	"nfxid/modules/directory/infrastructure/repository/user_preferences/delete"
	"nfxid/modules/directory/infrastructure/repository/user_preferences/get"
	"nfxid/modules/directory/infrastructure/repository/user_preferences/update"

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
