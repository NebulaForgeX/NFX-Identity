package grants

import (
	"nfxid/modules/access/domain/grants"
	"nfxid/modules/access/infrastructure/repository/grants/check"
	"nfxid/modules/access/infrastructure/repository/grants/create"
	"nfxid/modules/access/infrastructure/repository/grants/delete"
	"nfxid/modules/access/infrastructure/repository/grants/get"
	"nfxid/modules/access/infrastructure/repository/grants/update"

	"gorm.io/gorm"
)

// NewRepo 创建一个新的 Grant repository
func NewRepo(db *gorm.DB) *grants.Repo {
	return &grants.Repo{
		Create: create.NewHandler(db),
		Get:    get.NewHandler(db),
		Check:  check.NewHandler(db),
		Update: update.NewHandler(db),
		Delete: delete.NewHandler(db),
	}
}
