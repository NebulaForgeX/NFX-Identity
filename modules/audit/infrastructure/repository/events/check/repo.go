package check

import (
	"nfxid/modules/audit/domain/events"

	"gorm.io/gorm"
)

// Handler 处理检查操作，实现 events.Check 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Check Handler
func NewHandler(db *gorm.DB) events.Check {
	return &Handler{db: db}
}
