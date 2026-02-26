package user_educations

import (
	"nfxidentity/modules/directory/domain/user_educations"
	"nfxidentity/modules/directory/infrastructure/repository/user_educations/check"
	"nfxidentity/modules/directory/infrastructure/repository/user_educations/create"
	"nfxidentity/modules/directory/infrastructure/repository/user_educations/delete"
	"nfxidentity/modules/directory/infrastructure/repository/user_educations/get"
	"nfxidentity/modules/directory/infrastructure/repository/user_educations/update"

	"gorm.io/gorm"
)

// NewRepo 创建一个新的 UserEducation repository
func NewRepo(db *gorm.DB) *user_educations.Repo {
	return &user_educations.Repo{
		Create: create.NewHandler(db),
		Get:    get.NewHandler(db),
		Check:  check.NewHandler(db),
		Update: update.NewHandler(db),
		Delete: delete.NewHandler(db),
	}
}
