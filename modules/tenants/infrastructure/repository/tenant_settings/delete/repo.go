package delete

import (
	"nfxid/modules/tenants/domain/tenant_settings"

	"gorm.io/gorm"
)

// Handler 处理删除操作，实现 tenant_settings.Delete 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Delete Handler
func NewHandler(db *gorm.DB) tenant_settings.Delete {
	return &Handler{db: db}
}
