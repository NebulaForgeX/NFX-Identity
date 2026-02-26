package member_groups

import (
	"nfxidentity/modules/tenants/domain/member_groups"
	"nfxidentity/modules/tenants/infrastructure/repository/member_groups/check"
	"nfxidentity/modules/tenants/infrastructure/repository/member_groups/create"
	"nfxidentity/modules/tenants/infrastructure/repository/member_groups/delete"
	"nfxidentity/modules/tenants/infrastructure/repository/member_groups/get"
	"nfxidentity/modules/tenants/infrastructure/repository/member_groups/update"

	"gorm.io/gorm"
)

// NewRepo 创建一个新的 MemberGroup repository
func NewRepo(db *gorm.DB) *member_groups.Repo {
	return &member_groups.Repo{
		Create: create.NewHandler(db),
		Get:    get.NewHandler(db),
		Check:  check.NewHandler(db),
		Update: update.NewHandler(db),
		Delete: delete.NewHandler(db),
	}
}
