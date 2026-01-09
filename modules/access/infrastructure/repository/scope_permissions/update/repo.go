package update

import (
	"nfxid/modules/access/domain/scope_permissions"

	"gorm.io/gorm"
)

// Handler 处理更新操作，实现 scope_permissions.Update 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Update Handler
func NewHandler(db *gorm.DB) scope_permissions.Update {
	return &Handler{db: db}
}
