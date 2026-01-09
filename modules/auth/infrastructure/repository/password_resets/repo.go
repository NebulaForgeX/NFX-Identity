package password_resets

import (
	"nfxid/modules/auth/domain/password_resets"
	"nfxid/modules/auth/infrastructure/repository/password_resets/check"
	"nfxid/modules/auth/infrastructure/repository/password_resets/create"
	"nfxid/modules/auth/infrastructure/repository/password_resets/delete"
	"nfxid/modules/auth/infrastructure/repository/password_resets/get"
	"nfxid/modules/auth/infrastructure/repository/password_resets/update"

	"gorm.io/gorm"
)

// NewRepo 创建一个新的 PasswordReset repository
func NewRepo(db *gorm.DB) *password_resets.Repo {
	return &password_resets.Repo{
		Create: create.NewHandler(db),
		Get:    get.NewHandler(db),
		Check:  check.NewHandler(db),
		Update: update.NewHandler(db),
		Delete: delete.NewHandler(db),
	}
}
