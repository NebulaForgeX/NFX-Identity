package member_groups

import (
	"nfxid/modules/tenants/domain/member_groups"
	"nfxid/modules/tenants/infrastructure/repository/member_groups/check"
	"nfxid/modules/tenants/infrastructure/repository/member_groups/create"
	"nfxid/modules/tenants/infrastructure/repository/member_groups/delete"
	"nfxid/modules/tenants/infrastructure/repository/member_groups/get"
	"nfxid/modules/tenants/infrastructure/repository/member_groups/update"

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
