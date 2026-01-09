package invitations

import (
	"nfxid/modules/tenants/domain/invitations"
	"nfxid/modules/tenants/infrastructure/repository/invitations/check"
	"nfxid/modules/tenants/infrastructure/repository/invitations/create"
	"nfxid/modules/tenants/infrastructure/repository/invitations/delete"
	"nfxid/modules/tenants/infrastructure/repository/invitations/get"
	"nfxid/modules/tenants/infrastructure/repository/invitations/update"

	"gorm.io/gorm"
)

// NewRepo 创建一个新的 Invitation repository
func NewRepo(db *gorm.DB) *invitations.Repo {
	return &invitations.Repo{
		Create: create.NewHandler(db),
		Get:    get.NewHandler(db),
		Check:  check.NewHandler(db),
		Update: update.NewHandler(db),
		Delete: delete.NewHandler(db),
	}
}
