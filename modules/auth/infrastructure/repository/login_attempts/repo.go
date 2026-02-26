package login_attempts

import (
	"nfxidentity/modules/auth/domain/login_attempts"
	"nfxidentity/modules/auth/infrastructure/repository/login_attempts/check"
	"nfxidentity/modules/auth/infrastructure/repository/login_attempts/create"
	"nfxidentity/modules/auth/infrastructure/repository/login_attempts/delete"
	"nfxidentity/modules/auth/infrastructure/repository/login_attempts/get"
	"nfxidentity/modules/auth/infrastructure/repository/login_attempts/update"

	"gorm.io/gorm"
)

// NewRepo 创建一个新的 LoginAttempt repository
func NewRepo(db *gorm.DB) *login_attempts.Repo {
	return &login_attempts.Repo{
		Create: create.NewHandler(db),
		Get:    get.NewHandler(db),
		Check:  check.NewHandler(db),
		Update: update.NewHandler(db),
		Delete: delete.NewHandler(db),
	}
}
