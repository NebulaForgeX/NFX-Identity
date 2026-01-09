package get

import (
	"nfxid/modules/audit/domain/actor_snapshots"

	"gorm.io/gorm"
)

// Handler 处理获取数据操作，实现 actor_snapshots.Get 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Get Handler
func NewHandler(db *gorm.DB) actor_snapshots.Get {
	return &Handler{db: db}
}
