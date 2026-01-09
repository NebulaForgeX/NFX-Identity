package events

import (
	"nfxid/modules/audit/domain/events"
	"nfxid/modules/audit/infrastructure/repository/events/check"
	"nfxid/modules/audit/infrastructure/repository/events/create"
	"nfxid/modules/audit/infrastructure/repository/events/delete"
	"nfxid/modules/audit/infrastructure/repository/events/get"
	"nfxid/modules/audit/infrastructure/repository/events/update"

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
