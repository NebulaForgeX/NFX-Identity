package members

import (
	"nfxidentity/modules/tenants/domain/members"
	"nfxidentity/modules/tenants/infrastructure/repository/members/check"
	"nfxidentity/modules/tenants/infrastructure/repository/members/create"
	"nfxidentity/modules/tenants/infrastructure/repository/members/delete"
	"nfxidentity/modules/tenants/infrastructure/repository/members/get"
	"nfxidentity/modules/tenants/infrastructure/repository/members/update"

	"gorm.io/gorm"
)

// NewRepo 创建一个新的 Member repository
func NewRepo(db *gorm.DB) *members.Repo {
	return &members.Repo{
		Create: create.NewHandler(db),
		Get:    get.NewHandler(db),
		Check:  check.NewHandler(db),
		Update: update.NewHandler(db),
		Delete: delete.NewHandler(db),
	}
}
