package scopes

import (
	"nfxid/modules/access/domain/scopes"
	"nfxid/modules/access/infrastructure/repository/scopes/check"
	"nfxid/modules/access/infrastructure/repository/scopes/create"
	"nfxid/modules/access/infrastructure/repository/scopes/delete"
	"nfxid/modules/access/infrastructure/repository/scopes/get"
	"nfxid/modules/access/infrastructure/repository/scopes/update"

	"gorm.io/gorm"
)

// NewRepo 创建一个新的 Scope repository
func NewRepo(db *gorm.DB) *scopes.Repo {
	return &scopes.Repo{
		Create: create.NewHandler(db),
		Get:    get.NewHandler(db),
		Check:  check.NewHandler(db),
		Update: update.NewHandler(db),
		Delete: delete.NewHandler(db),
	}
}
