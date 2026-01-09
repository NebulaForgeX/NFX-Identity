package hash_chain_checkpoints

import (
	"nfxid/modules/audit/domain/hash_chain_checkpoints"
	"nfxid/modules/audit/infrastructure/repository/hash_chain_checkpoints/check"
	"nfxid/modules/audit/infrastructure/repository/hash_chain_checkpoints/create"
	"nfxid/modules/audit/infrastructure/repository/hash_chain_checkpoints/delete"
	"nfxid/modules/audit/infrastructure/repository/hash_chain_checkpoints/get"
	"nfxid/modules/audit/infrastructure/repository/hash_chain_checkpoints/update"

	"gorm.io/gorm"
)

// NewRepo 创建一个新的 HashChainCheckpoint repository
func NewRepo(db *gorm.DB) *hash_chain_checkpoints.Repo {
	return &hash_chain_checkpoints.Repo{
		Create: create.NewHandler(db),
		Get:    get.NewHandler(db),
		Check:  check.NewHandler(db),
		Update: update.NewHandler(db),
		Delete: delete.NewHandler(db),
	}
}
