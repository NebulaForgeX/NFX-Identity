package password_history

import (
	"nfxid/modules/auth/domain/password_history"
	"nfxid/modules/auth/infrastructure/repository/password_history/check"
	"nfxid/modules/auth/infrastructure/repository/password_history/create"
	"nfxid/modules/auth/infrastructure/repository/password_history/delete"
	"nfxid/modules/auth/infrastructure/repository/password_history/get"
	"nfxid/modules/auth/infrastructure/repository/password_history/update"

	"gorm.io/gorm"
)

// NewRepo 创建一个新的 PasswordHistory repository
func NewRepo(db *gorm.DB) *password_history.Repo {
	return &password_history.Repo{
		Create: create.NewHandler(db),
		Get:    get.NewHandler(db),
		Check:  check.NewHandler(db),
		Update: update.NewHandler(db),
		Delete: delete.NewHandler(db),
	}
}
