package update

import (
	"nfxid/modules/audit/domain/event_search_index"

	"gorm.io/gorm"
)

// Handler 处理更新操作，实现 event_search_index.Update 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Update Handler
func NewHandler(db *gorm.DB) event_search_index.Update {
	return &Handler{db: db}
}
