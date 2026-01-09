package update

import (
	"nfxid/modules/audit/domain/event_retention_policies"

	"gorm.io/gorm"
)

// Handler 处理更新操作，实现 event_retention_policies.Update 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Update Handler
func NewHandler(db *gorm.DB) event_retention_policies.Update {
	return &Handler{db: db}
}
