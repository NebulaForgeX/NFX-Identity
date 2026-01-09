package get

import (
	"nfxid/modules/access/domain/scope_permissions"

	"gorm.io/gorm"
)

// Handler 处理获取数据操作，实现 scope_permissions.Get 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Get Handler
func NewHandler(db *gorm.DB) scope_permissions.Get {
	return &Handler{db: db}
}
