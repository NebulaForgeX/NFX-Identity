package update

import (
	"nfxid/modules/audit/domain/events"

	"gorm.io/gorm"
)

// Handler 处理更新操作，实现 events.Update 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Update Handler
func NewHandler(db *gorm.DB) events.Update {
	return &Handler{db: db}
}
