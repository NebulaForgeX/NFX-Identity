package application_roles

import (
	dom "nfxid/modules/access/domain/application_roles"
	"nfxid/modules/access/infrastructure/repository/application_roles/check"
	"nfxid/modules/access/infrastructure/repository/application_roles/create"
	"nfxid/modules/access/infrastructure/repository/application_roles/delete"
	"nfxid/modules/access/infrastructure/repository/application_roles/get"
	"nfxid/modules/access/infrastructure/repository/application_roles/update"
	"gorm.io/gorm"
)

func NewRepo(db *gorm.DB) *dom.Repo {
	return &dom.Repo{
		Create: create.NewHandler(db),
		Get:    get.NewHandler(db),
		Check:  check.NewHandler(db),
		Update: update.NewHandler(db),
		Delete: delete.NewHandler(db),
	}
}
