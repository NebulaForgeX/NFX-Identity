package roles

import (
	"nfxid/modules/access/domain/roles"
	"nfxid/modules/access/infrastructure/repository/roles/check"
	"nfxid/modules/access/infrastructure/repository/roles/create"
	"nfxid/modules/access/infrastructure/repository/roles/delete"
	"nfxid/modules/access/infrastructure/repository/roles/get"
	"nfxid/modules/access/infrastructure/repository/roles/update"

	"gorm.io/gorm"
)

// NewRepo 创建一个新的 Role repository
func NewRepo(db *gorm.DB) *roles.Repo {
	return &roles.Repo{
		Create: create.NewHandler(db),
		Get:    get.NewHandler(db),
		Check:  check.NewHandler(db),
		Update: update.NewHandler(db),
		Delete: delete.NewHandler(db),
	}
}
