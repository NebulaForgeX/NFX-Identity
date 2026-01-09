package delete

import (
	"nfxid/modules/audit/domain/actor_snapshots"

	"gorm.io/gorm"
)

// Handler 处理删除操作，实现 actor_snapshots.Delete 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Delete Handler
func NewHandler(db *gorm.DB) actor_snapshots.Delete {
	return &Handler{db: db}
}
