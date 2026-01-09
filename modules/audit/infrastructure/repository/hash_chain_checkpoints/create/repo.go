package create

import (
	"nfxid/modules/audit/domain/hash_chain_checkpoints"

	"gorm.io/gorm"
)

// Handler 处理创建操作，实现 hash_chain_checkpoints.Create 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Create Handler
func NewHandler(db *gorm.DB) hash_chain_checkpoints.Create {
	return &Handler{db: db}
}
