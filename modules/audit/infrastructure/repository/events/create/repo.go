package create

import (
	"nfxid/modules/audit/domain/events"

	"gorm.io/gorm"
)

// Handler 处理创建操作，实现 events.Create 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Create Handler
func NewHandler(db *gorm.DB) events.Create {
	return &Handler{db: db}
}
