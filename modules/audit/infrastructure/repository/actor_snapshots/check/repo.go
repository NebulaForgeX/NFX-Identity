package check

import (
	"nfxid/modules/audit/domain/actor_snapshots"

	"gorm.io/gorm"
)

// Handler 处理检查操作，实现 actor_snapshots.Check 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Check Handler
func NewHandler(db *gorm.DB) actor_snapshots.Check {
	return &Handler{db: db}
}
