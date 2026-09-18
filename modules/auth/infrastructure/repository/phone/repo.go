package phone

import (
	"nfxidentity/modules/auth/domain/phone"
	"nfxidentity/modules/auth/infrastructure/repository/phone/create"
	"nfxidentity/modules/auth/infrastructure/repository/phone/get"
	"nfxidentity/modules/auth/infrastructure/repository/phone/update"

	"gorm.io/gorm"
)

func NewRepo(db *gorm.DB) *phone.Repo {
	return &phone.Repo{Create: create.NewHandler(db), Get: get.NewHandler(db), Update: update.NewHandler(db)}
}
