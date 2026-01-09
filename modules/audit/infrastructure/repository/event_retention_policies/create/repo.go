package create

import (
	"nfxid/modules/audit/domain/event_retention_policies"

	"gorm.io/gorm"
)

// Handler 处理创建操作，实现 event_retention_policies.Create 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Create Handler
func NewHandler(db *gorm.DB) event_retention_policies.Create {
	return &Handler{db: db}
}
