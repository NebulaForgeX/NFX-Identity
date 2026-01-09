package update

import (
	"nfxid/modules/audit/domain/hash_chain_checkpoints"

	"gorm.io/gorm"
)

// Handler 处理更新操作，实现 hash_chain_checkpoints.Update 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Update Handler
func NewHandler(db *gorm.DB) hash_chain_checkpoints.Update {
	return &Handler{db: db}
}
