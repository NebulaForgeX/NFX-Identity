package create

import (
	"nfxid/modules/access/domain/permissions"

	"gorm.io/gorm"
)

// Handler 处理创建操作，实现 permissions.Create 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Create Handler
func NewHandler(db *gorm.DB) permissions.Create {
	return &Handler{db: db}
}
