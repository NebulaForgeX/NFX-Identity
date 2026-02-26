package member_roles

import (
	"nfxidentity/modules/tenants/domain/member_roles"
	"nfxidentity/modules/tenants/infrastructure/repository/member_roles/check"
	"nfxidentity/modules/tenants/infrastructure/repository/member_roles/create"
	"nfxidentity/modules/tenants/infrastructure/repository/member_roles/delete"
	"nfxidentity/modules/tenants/infrastructure/repository/member_roles/get"
	"nfxidentity/modules/tenants/infrastructure/repository/member_roles/update"

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
