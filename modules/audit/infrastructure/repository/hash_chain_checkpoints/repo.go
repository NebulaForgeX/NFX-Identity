package hash_chain_checkpoints

import (
	"nfxidentity/modules/audit/domain/hash_chain_checkpoints"
	"nfxidentity/modules/audit/infrastructure/repository/hash_chain_checkpoints/check"
	"nfxidentity/modules/audit/infrastructure/repository/hash_chain_checkpoints/create"
	"nfxidentity/modules/audit/infrastructure/repository/hash_chain_checkpoints/delete"
	"nfxidentity/modules/audit/infrastructure/repository/hash_chain_checkpoints/get"
	"nfxidentity/modules/audit/infrastructure/repository/hash_chain_checkpoints/update"

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
