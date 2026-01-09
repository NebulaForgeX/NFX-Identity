package groups

import (
	"nfxid/modules/tenants/domain/groups"
	"nfxid/modules/tenants/infrastructure/repository/groups/check"
	"nfxid/modules/tenants/infrastructure/repository/groups/create"
	"nfxid/modules/tenants/infrastructure/repository/groups/delete"
	"nfxid/modules/tenants/infrastructure/repository/groups/get"
	"nfxid/modules/tenants/infrastructure/repository/groups/update"

	"gorm.io/gorm"
)

// NewRepo 创建一个新的 Group repository
func NewRepo(db *gorm.DB) *groups.Repo {
	return &groups.Repo{
		Create: create.NewHandler(db),
		Get:    get.NewHandler(db),
		Check:  check.NewHandler(db),
		Update: update.NewHandler(db),
		Delete: delete.NewHandler(db),
	}
}
