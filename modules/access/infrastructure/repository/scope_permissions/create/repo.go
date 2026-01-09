package create

import (
	"nfxid/modules/access/domain/scope_permissions"

	"gorm.io/gorm"
)

// Handler 处理创建操作，实现 scope_permissions.Create 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Create Handler
func NewHandler(db *gorm.DB) scope_permissions.Create {
	return &Handler{db: db}
}
