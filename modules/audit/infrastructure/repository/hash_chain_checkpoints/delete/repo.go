package delete

import (
	"nfxid/modules/audit/domain/hash_chain_checkpoints"

	"gorm.io/gorm"
)

// Handler 处理删除操作，实现 hash_chain_checkpoints.Delete 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Delete Handler
func NewHandler(db *gorm.DB) hash_chain_checkpoints.Delete {
	return &Handler{db: db}
}
