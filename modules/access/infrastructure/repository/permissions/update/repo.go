package update

import (
	"nfxid/modules/access/domain/permissions"

	"gorm.io/gorm"
)

// Handler 处理更新操作，实现 permissions.Update 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Update Handler
func NewHandler(db *gorm.DB) permissions.Update {
	return &Handler{db: db}
}
