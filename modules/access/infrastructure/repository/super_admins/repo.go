package super_admins

import (
	"nfxid/modules/access/domain/super_admins"
	"nfxid/modules/access/infrastructure/repository/super_admins/check"
	"nfxid/modules/access/infrastructure/repository/super_admins/create"
	"nfxid/modules/access/infrastructure/repository/super_admins/delete"
	"nfxid/modules/access/infrastructure/repository/super_admins/get"
	"gorm.io/gorm"
)

func NewRepo(db *gorm.DB) *super_admins.Repo {
	return &super_admins.Repo{
		Create: create.NewHandler(db),
		Get:    get.NewHandler(db),
		Check:  check.NewHandler(db),
		Delete: delete.NewHandler(db),
	}
}
