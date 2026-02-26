package actor_snapshots

import (
	"nfxidentity/modules/audit/domain/actor_snapshots"
	"nfxidentity/modules/audit/infrastructure/repository/actor_snapshots/check"
	"nfxidentity/modules/audit/infrastructure/repository/actor_snapshots/create"
	"nfxidentity/modules/audit/infrastructure/repository/actor_snapshots/delete"
	"nfxidentity/modules/audit/infrastructure/repository/actor_snapshots/get"
	"nfxidentity/modules/audit/infrastructure/repository/actor_snapshots/update"

	"gorm.io/gorm"
)

// NewRepo 创建一个新的 ActorSnapshot repository
func NewRepo(db *gorm.DB) *actor_snapshots.Repo {
	return &actor_snapshots.Repo{
		Create: create.NewHandler(db),
		Get:    get.NewHandler(db),
		Check:  check.NewHandler(db),
		Update: update.NewHandler(db),
		Delete: delete.NewHandler(db),
	}
}
