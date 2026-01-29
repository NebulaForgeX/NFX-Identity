package action_requirements

import (
	"nfxid/modules/access/domain/action_requirements"
	"nfxid/modules/access/infrastructure/repository/action_requirements/check"
	"nfxid/modules/access/infrastructure/repository/action_requirements/create"
	"nfxid/modules/access/infrastructure/repository/action_requirements/delete"
	"nfxid/modules/access/infrastructure/repository/action_requirements/get"

	"gorm.io/gorm"
)

func NewRepo(db *gorm.DB) *action_requirements.Repo {
	return &action_requirements.Repo{
		Create: create.NewHandler(db),
		Get:    get.NewHandler(db),
		Check:  check.NewHandler(db),
		Delete: delete.NewHandler(db),
	}
}
