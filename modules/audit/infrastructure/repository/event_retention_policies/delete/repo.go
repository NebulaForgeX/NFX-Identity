package delete

import (
	"nfxid/modules/audit/domain/event_retention_policies"

	"gorm.io/gorm"
)

// Handler 处理删除操作，实现 event_retention_policies.Delete 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Delete Handler
func NewHandler(db *gorm.DB) event_retention_policies.Delete {
	return &Handler{db: db}
}
