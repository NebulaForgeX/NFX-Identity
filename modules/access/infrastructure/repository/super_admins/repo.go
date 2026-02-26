package super_admins

import (
	"nfxidentity/modules/access/domain/super_admins"
	"nfxidentity/modules/access/infrastructure/repository/super_admins/check"
	"nfxidentity/modules/access/infrastructure/repository/super_admins/create"
	"nfxidentity/modules/access/infrastructure/repository/super_admins/delete"
	"nfxidentity/modules/access/infrastructure/repository/super_admins/get"

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
