package update

import (
	"nfxid/modules/clients/domain/rate_limits"

	"gorm.io/gorm"
)

// Handler 处理更新操作，实现 rate_limits.Update 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Update Handler
func NewHandler(db *gorm.DB) rate_limits.Update {
	return &Handler{db: db}
}
