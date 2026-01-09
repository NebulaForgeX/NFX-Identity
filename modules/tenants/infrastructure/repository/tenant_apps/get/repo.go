package get

import (
	"nfxid/modules/tenants/domain/tenant_apps"

	"gorm.io/gorm"
)

// Handler 处理获取数据操作，实现 tenant_apps.Get 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Get Handler
func NewHandler(db *gorm.DB) tenant_apps.Get {
	return &Handler{db: db}
}
