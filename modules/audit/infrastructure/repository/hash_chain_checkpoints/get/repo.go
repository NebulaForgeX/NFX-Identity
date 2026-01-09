package get

import (
	"nfxid/modules/audit/domain/hash_chain_checkpoints"

	"gorm.io/gorm"
)

// Handler 处理获取数据操作，实现 hash_chain_checkpoints.Get 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Get Handler
func NewHandler(db *gorm.DB) hash_chain_checkpoints.Get {
	return &Handler{db: db}
}
