package check

import (
	"nfxid/modules/audit/domain/event_retention_policies"

	"gorm.io/gorm"
)

// Handler 处理检查操作，实现 event_retention_policies.Check 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Check Handler
func NewHandler(db *gorm.DB) event_retention_policies.Check {
	return &Handler{db: db}
}
