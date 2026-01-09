package update

import (
	"nfxid/modules/audit/domain/actor_snapshots"

	"gorm.io/gorm"
)

// Handler 处理更新操作，实现 actor_snapshots.Update 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Update Handler
func NewHandler(db *gorm.DB) actor_snapshots.Update {
	return &Handler{db: db}
}
