package check

import (
	"nfxid/modules/audit/domain/event_search_index"

	"gorm.io/gorm"
)

// Handler 处理检查操作，实现 event_search_index.Check 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Check Handler
func NewHandler(db *gorm.DB) event_search_index.Check {
	return &Handler{db: db}
}
