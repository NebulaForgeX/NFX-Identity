package role

import (
	"nfxid/modules/auth/domain/role"
	"nfxid/modules/auth/infrastructure/repository/role/check"
	"nfxid/modules/auth/infrastructure/repository/role/create"
	"nfxid/modules/auth/infrastructure/repository/role/delete"
	"nfxid/modules/auth/infrastructure/repository/role/get"
	"nfxid/modules/auth/infrastructure/repository/role/update"

	"gorm.io/gorm"
)

// NewRepo 创建一个新的 Role repository
func NewRepo(db *gorm.DB) *role.Repo {
	return &role.Repo{
		Create: create.NewHandler(db),
		Get:    get.NewHandler(db),
		Check:  check.NewHandler(db),
		Update: update.NewHandler(db),
		Delete: delete.NewHandler(db),
	}
}
