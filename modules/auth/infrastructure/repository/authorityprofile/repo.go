package authorityprofile

import (
	"nfxidentity/modules/auth/domain/authorityprofile"
	"nfxidentity/modules/auth/infrastructure/repository/authorityprofile/create"
	"nfxidentity/modules/auth/infrastructure/repository/authorityprofile/get"
	"nfxidentity/modules/auth/infrastructure/repository/authorityprofile/update"

	"gorm.io/gorm"
)

func NewRepo(db *gorm.DB) *authorityprofile.Repo {
	return &authorityprofile.Repo{Create: create.NewHandler(db), Get: get.NewHandler(db), Update: update.NewHandler(db)}
}
