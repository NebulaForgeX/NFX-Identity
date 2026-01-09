package delete

import (
	"nfxid/modules/audit/domain/events"

	"gorm.io/gorm"
)

// Handler 处理删除操作，实现 events.Delete 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Delete Handler
func NewHandler(db *gorm.DB) events.Delete {
	return &Handler{db: db}
}
