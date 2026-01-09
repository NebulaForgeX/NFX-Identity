package delete

import (
	"nfxid/modules/clients/domain/rate_limits"

	"gorm.io/gorm"
)

// Handler 处理删除操作，实现 rate_limits.Delete 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Delete Handler
func NewHandler(db *gorm.DB) rate_limits.Delete {
	return &Handler{db: db}
}
