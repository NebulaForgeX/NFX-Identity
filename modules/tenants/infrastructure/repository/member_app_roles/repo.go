package member_app_roles

import (
	"nfxidentity/modules/tenants/domain/member_app_roles"
	"nfxidentity/modules/tenants/infrastructure/repository/member_app_roles/check"
	"nfxidentity/modules/tenants/infrastructure/repository/member_app_roles/create"
	"nfxidentity/modules/tenants/infrastructure/repository/member_app_roles/delete"
	"nfxidentity/modules/tenants/infrastructure/repository/member_app_roles/get"
	"nfxidentity/modules/tenants/infrastructure/repository/member_app_roles/update"

	"gorm.io/gorm"
)

// NewRepo 创建一个新的 MemberAppRole repository
func NewRepo(db *gorm.DB) *member_app_roles.Repo {
	return &member_app_roles.Repo{
		Create: create.NewHandler(db),
		Get:    get.NewHandler(db),
		Check:  check.NewHandler(db),
		Update: update.NewHandler(db),
		Delete: delete.NewHandler(db),
	}
}
