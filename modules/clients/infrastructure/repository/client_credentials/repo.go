package client_credentials

import (
	"nfxidentity/modules/clients/domain/client_credentials"
	"nfxidentity/modules/clients/infrastructure/repository/client_credentials/check"
	"nfxidentity/modules/clients/infrastructure/repository/client_credentials/create"
	"nfxidentity/modules/clients/infrastructure/repository/client_credentials/delete"
	"nfxidentity/modules/clients/infrastructure/repository/client_credentials/get"
	"nfxidentity/modules/clients/infrastructure/repository/client_credentials/update"

	"gorm.io/gorm"
)

// NewRepo 创建一个新的 ClientCredential repository
func NewRepo(db *gorm.DB) *client_credentials.Repo {
	return &client_credentials.Repo{
		Create: create.NewHandler(db),
		Get:    get.NewHandler(db),
		Check:  check.NewHandler(db),
		Update: update.NewHandler(db),
		Delete: delete.NewHandler(db),
	}
}
