package events

import (
	"nfxidentity/modules/audit/domain/events"
	"nfxidentity/modules/audit/infrastructure/repository/events/check"
	"nfxidentity/modules/audit/infrastructure/repository/events/create"
	"nfxidentity/modules/audit/infrastructure/repository/events/delete"
	"nfxidentity/modules/audit/infrastructure/repository/events/get"
	"nfxidentity/modules/audit/infrastructure/repository/events/update"

	"gorm.io/gorm"
)

// NewRepo 创建一个新的 Event repository
func NewRepo(db *gorm.DB) *events.Repo {
	return &events.Repo{
		Create: create.NewHandler(db),
		Get:    get.NewHandler(db),
		Check:  check.NewHandler(db),
		Update: update.NewHandler(db),
		Delete: delete.NewHandler(db),
	}
}
