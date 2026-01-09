package create

import (
	"nfxid/modules/tenants/domain/tenants"

	"gorm.io/gorm"
)

// Handler 处理创建操作，实现 tenants.Create 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Create Handler
func NewHandler(db *gorm.DB) tenants.Create {
	return &Handler{db: db}
}
