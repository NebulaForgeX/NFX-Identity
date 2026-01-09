package create

import (
	"nfxid/modules/tenants/domain/tenant_apps"

	"gorm.io/gorm"
)

// Handler 处理创建操作，实现 tenant_apps.Create 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Create Handler
func NewHandler(db *gorm.DB) tenant_apps.Create {
	return &Handler{db: db}
}
