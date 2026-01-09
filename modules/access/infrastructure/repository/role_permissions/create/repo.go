package create

import (
	"nfxid/modules/access/domain/role_permissions"

	"gorm.io/gorm"
)

// Handler 处理创建操作，实现 role_permissions.Create 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Create Handler
func NewHandler(db *gorm.DB) role_permissions.Create {
	return &Handler{db: db}
}
