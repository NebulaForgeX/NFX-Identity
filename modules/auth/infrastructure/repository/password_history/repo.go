package password_history

import (
	"nfxidentity/modules/auth/domain/password_history"
	"nfxidentity/modules/auth/infrastructure/repository/password_history/check"
	"nfxidentity/modules/auth/infrastructure/repository/password_history/create"
	"nfxidentity/modules/auth/infrastructure/repository/password_history/delete"
	"nfxidentity/modules/auth/infrastructure/repository/password_history/get"
	"nfxidentity/modules/auth/infrastructure/repository/password_history/update"

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
