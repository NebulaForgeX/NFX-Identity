package member_roles

import (
	"nfxid/modules/tenants/domain/member_roles"
	"nfxid/modules/tenants/infrastructure/repository/member_roles/check"
	"nfxid/modules/tenants/infrastructure/repository/member_roles/create"
	"nfxid/modules/tenants/infrastructure/repository/member_roles/delete"
	"nfxid/modules/tenants/infrastructure/repository/member_roles/get"
	"nfxid/modules/tenants/infrastructure/repository/member_roles/update"

	"gorm.io/gorm"
)

// NewRepo 创建一个新的 MemberRole repository
func NewRepo(db *gorm.DB) *member_roles.Repo {
	return &member_roles.Repo{
		Create: create.NewHandler(db),
		Get:    get.NewHandler(db),
		Check:  check.NewHandler(db),
		Update: update.NewHandler(db),
		Delete: delete.NewHandler(db),
	}
}
