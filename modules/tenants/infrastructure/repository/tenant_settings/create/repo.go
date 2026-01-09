package create

import (
	"nfxid/modules/tenants/domain/tenant_settings"

	"gorm.io/gorm"
)

// Handler 处理创建操作，实现 tenant_settings.Create 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Create Handler
func NewHandler(db *gorm.DB) tenant_settings.Create {
	return &Handler{db: db}
}
