package event_retention_policies

import (
	"nfxid/modules/audit/domain/event_retention_policies"
	"nfxid/modules/audit/infrastructure/repository/event_retention_policies/check"
	"nfxid/modules/audit/infrastructure/repository/event_retention_policies/create"
	"nfxid/modules/audit/infrastructure/repository/event_retention_policies/delete"
	"nfxid/modules/audit/infrastructure/repository/event_retention_policies/get"
	"nfxid/modules/audit/infrastructure/repository/event_retention_policies/update"

	"gorm.io/gorm"
)

// NewRepo 创建一个新的 EventRetentionPolicy repository
func NewRepo(db *gorm.DB) *event_retention_policies.Repo {
	return &event_retention_policies.Repo{
		Create: create.NewHandler(db),
		Get:    get.NewHandler(db),
		Check:  check.NewHandler(db),
		Update: update.NewHandler(db),
		Delete: delete.NewHandler(db),
	}
}
