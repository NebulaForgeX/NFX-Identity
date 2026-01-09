package login_attempts

import (
	"nfxid/modules/auth/domain/login_attempts"
	"nfxid/modules/auth/infrastructure/repository/login_attempts/check"
	"nfxid/modules/auth/infrastructure/repository/login_attempts/create"
	"nfxid/modules/auth/infrastructure/repository/login_attempts/delete"
	"nfxid/modules/auth/infrastructure/repository/login_attempts/get"
	"nfxid/modules/auth/infrastructure/repository/login_attempts/update"

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
