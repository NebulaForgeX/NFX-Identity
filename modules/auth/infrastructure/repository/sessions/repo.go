package sessions

import (
	"nfxid/modules/auth/domain/sessions"
	"nfxid/modules/auth/infrastructure/repository/sessions/check"
	"nfxid/modules/auth/infrastructure/repository/sessions/create"
	"nfxid/modules/auth/infrastructure/repository/sessions/delete"
	"nfxid/modules/auth/infrastructure/repository/sessions/get"
	"nfxid/modules/auth/infrastructure/repository/sessions/update"

	"gorm.io/gorm"
)

// NewRepo 创建一个新的 Session repository
func NewRepo(db *gorm.DB) *sessions.Repo {
	return &sessions.Repo{
		Create: create.NewHandler(db),
		Get:    get.NewHandler(db),
		Check:  check.NewHandler(db),
		Update: update.NewHandler(db),
		Delete: delete.NewHandler(db),
	}
}
