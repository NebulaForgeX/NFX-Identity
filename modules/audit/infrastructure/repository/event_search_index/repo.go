package event_search_index

import (
	"nfxid/modules/audit/domain/event_search_index"
	"nfxid/modules/audit/infrastructure/repository/event_search_index/check"
	"nfxid/modules/audit/infrastructure/repository/event_search_index/create"
	"nfxid/modules/audit/infrastructure/repository/event_search_index/delete"
	"nfxid/modules/audit/infrastructure/repository/event_search_index/get"
	"nfxid/modules/audit/infrastructure/repository/event_search_index/update"

	"gorm.io/gorm"
)

// NewRepo 创建一个新的 EventSearchIndex repository
func NewRepo(db *gorm.DB) *event_search_index.Repo {
	return &event_search_index.Repo{
		Create: create.NewHandler(db),
		Get:    get.NewHandler(db),
		Check:  check.NewHandler(db),
		Update: update.NewHandler(db),
		Delete: delete.NewHandler(db),
	}
}
