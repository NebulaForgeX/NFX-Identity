package create

import (
	"nfxid/modules/audit/domain/event_search_index"

	"gorm.io/gorm"
)

// Handler 处理创建操作，实现 event_search_index.Create 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Create Handler
func NewHandler(db *gorm.DB) event_search_index.Create {
	return &Handler{db: db}
}
