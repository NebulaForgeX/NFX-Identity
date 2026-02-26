package sessions

import (
	"nfxidentity/modules/auth/domain/sessions"
	"nfxidentity/modules/auth/infrastructure/repository/sessions/check"
	"nfxidentity/modules/auth/infrastructure/repository/sessions/create"
	"nfxidentity/modules/auth/infrastructure/repository/sessions/delete"
	"nfxidentity/modules/auth/infrastructure/repository/sessions/get"
	"nfxidentity/modules/auth/infrastructure/repository/sessions/update"

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
