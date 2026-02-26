package invitations

import (
	"nfxidentity/modules/tenants/domain/invitations"
	"nfxidentity/modules/tenants/infrastructure/repository/invitations/check"
	"nfxidentity/modules/tenants/infrastructure/repository/invitations/create"
	"nfxidentity/modules/tenants/infrastructure/repository/invitations/delete"
	"nfxidentity/modules/tenants/infrastructure/repository/invitations/get"
	"nfxidentity/modules/tenants/infrastructure/repository/invitations/update"

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
