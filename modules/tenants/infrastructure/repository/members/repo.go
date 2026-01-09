package members

import (
	"nfxid/modules/tenants/domain/members"
	"nfxid/modules/tenants/infrastructure/repository/members/check"
	"nfxid/modules/tenants/infrastructure/repository/members/create"
	"nfxid/modules/tenants/infrastructure/repository/members/delete"
	"nfxid/modules/tenants/infrastructure/repository/members/get"
	"nfxid/modules/tenants/infrastructure/repository/members/update"

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
