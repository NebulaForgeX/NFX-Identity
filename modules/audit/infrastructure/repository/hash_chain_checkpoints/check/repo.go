package check

import (
	"nfxid/modules/audit/domain/hash_chain_checkpoints"

	"gorm.io/gorm"
)

// Handler 处理检查操作，实现 hash_chain_checkpoints.Check 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Check Handler
func NewHandler(db *gorm.DB) hash_chain_checkpoints.Check {
	return &Handler{db: db}
}
