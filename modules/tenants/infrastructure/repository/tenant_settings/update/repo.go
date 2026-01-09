package update

import (
	"nfxid/modules/tenants/domain/tenant_settings"

	"gorm.io/gorm"
)

// Handler 处理更新操作，实现 tenant_settings.Update 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Update Handler
func NewHandler(db *gorm.DB) tenant_settings.Update {
	return &Handler{db: db}
}
