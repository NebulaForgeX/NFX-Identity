package create

import (
	"nfxid/modules/access/domain/roles"

	"gorm.io/gorm"
)

// Handler 处理创建操作，实现 roles.Create 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Create Handler
func NewHandler(db *gorm.DB) roles.Create {
	return &Handler{db: db}
}
