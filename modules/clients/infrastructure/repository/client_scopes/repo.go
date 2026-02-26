package client_scopes

import (
	"nfxidentity/modules/clients/domain/client_scopes"
	"nfxidentity/modules/clients/infrastructure/repository/client_scopes/check"
	"nfxidentity/modules/clients/infrastructure/repository/client_scopes/create"
	"nfxidentity/modules/clients/infrastructure/repository/client_scopes/delete"
	"nfxidentity/modules/clients/infrastructure/repository/client_scopes/get"
	"nfxidentity/modules/clients/infrastructure/repository/client_scopes/update"

	"gorm.io/gorm"
)

// NewRepo 创建一个新的 ClientScope repository
func NewRepo(db *gorm.DB) *client_scopes.Repo {
	return &client_scopes.Repo{
		Create: create.NewHandler(db),
		Get:    get.NewHandler(db),
		Check:  check.NewHandler(db),
		Update: update.NewHandler(db),
		Delete: delete.NewHandler(db),
	}
}
