package actions

import (
	"nfxid/modules/access/domain/actions"
	"nfxid/modules/access/infrastructure/repository/actions/check"
	"nfxid/modules/access/infrastructure/repository/actions/create"
	"nfxid/modules/access/infrastructure/repository/actions/delete"
	"nfxid/modules/access/infrastructure/repository/actions/get"
	"nfxid/modules/access/infrastructure/repository/actions/update"

	"gorm.io/gorm"
)

func NewRepo(db *gorm.DB) *actions.Repo {
	return &actions.Repo{
		Create: create.NewHandler(db),
		Get:    get.NewHandler(db),
		Check:  check.NewHandler(db),
		Update: update.NewHandler(db),
		Delete: delete.NewHandler(db),
	}
}
