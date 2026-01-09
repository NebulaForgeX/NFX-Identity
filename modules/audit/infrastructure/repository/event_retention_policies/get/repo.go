package get

import (
	"nfxid/modules/audit/domain/event_retention_policies"

	"gorm.io/gorm"
)

// Handler 处理获取数据操作，实现 event_retention_policies.Get 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Get Handler
func NewHandler(db *gorm.DB) event_retention_policies.Get {
	return &Handler{db: db}
}
