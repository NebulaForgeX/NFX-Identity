package event_retention_policies

import (
	"nfxidentity/modules/audit/domain/event_retention_policies"
	"nfxidentity/modules/audit/infrastructure/repository/event_retention_policies/check"
	"nfxidentity/modules/audit/infrastructure/repository/event_retention_policies/create"
	"nfxidentity/modules/audit/infrastructure/repository/event_retention_policies/delete"
	"nfxidentity/modules/audit/infrastructure/repository/event_retention_policies/get"
	"nfxidentity/modules/audit/infrastructure/repository/event_retention_policies/update"

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
