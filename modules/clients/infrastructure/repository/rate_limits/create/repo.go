package create

import (
	"nfxid/modules/clients/domain/rate_limits"

	"gorm.io/gorm"
)

// Handler 处理创建操作，实现 rate_limits.Create 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Create Handler
func NewHandler(db *gorm.DB) rate_limits.Create {
	return &Handler{db: db}
}
