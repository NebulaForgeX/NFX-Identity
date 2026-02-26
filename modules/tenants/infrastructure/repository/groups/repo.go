package groups

import (
	"nfxidentity/modules/tenants/domain/groups"
	"nfxidentity/modules/tenants/infrastructure/repository/groups/check"
	"nfxidentity/modules/tenants/infrastructure/repository/groups/create"
	"nfxidentity/modules/tenants/infrastructure/repository/groups/delete"
	"nfxidentity/modules/tenants/infrastructure/repository/groups/get"
	"nfxidentity/modules/tenants/infrastructure/repository/groups/update"

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
