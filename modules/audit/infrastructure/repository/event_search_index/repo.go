package event_search_index

import (
	"nfxidentity/modules/audit/domain/event_search_index"
	"nfxidentity/modules/audit/infrastructure/repository/event_search_index/check"
	"nfxidentity/modules/audit/infrastructure/repository/event_search_index/create"
	"nfxidentity/modules/audit/infrastructure/repository/event_search_index/delete"
	"nfxidentity/modules/audit/infrastructure/repository/event_search_index/get"
	"nfxidentity/modules/audit/infrastructure/repository/event_search_index/update"

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
