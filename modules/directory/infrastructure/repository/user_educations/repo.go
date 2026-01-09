package user_educations

import (
	"nfxid/modules/directory/domain/user_educations"
	"nfxid/modules/directory/infrastructure/repository/user_educations/check"
	"nfxid/modules/directory/infrastructure/repository/user_educations/create"
	"nfxid/modules/directory/infrastructure/repository/user_educations/delete"
	"nfxid/modules/directory/infrastructure/repository/user_educations/get"
	"nfxid/modules/directory/infrastructure/repository/user_educations/update"

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
