package check

import (
	"nfxid/modules/tenants/domain/tenant_apps"

	"gorm.io/gorm"
)

// Handler 处理检查操作，实现 tenant_apps.Check 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Check Handler
func NewHandler(db *gorm.DB) tenant_apps.Check {
	return &Handler{db: db}
}
