package actor_snapshots

import (
	"nfxid/modules/audit/domain/actor_snapshots"
	"nfxid/modules/audit/infrastructure/repository/actor_snapshots/check"
	"nfxid/modules/audit/infrastructure/repository/actor_snapshots/create"
	"nfxid/modules/audit/infrastructure/repository/actor_snapshots/delete"
	"nfxid/modules/audit/infrastructure/repository/actor_snapshots/get"
	"nfxid/modules/audit/infrastructure/repository/actor_snapshots/update"

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
