package check

import (
	"nfxid/modules/tenants/domain/tenants"

	"gorm.io/gorm"
)

// Handler 处理检查操作，实现 tenants.Check 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Check Handler
func NewHandler(db *gorm.DB) tenants.Check {
	return &Handler{db: db}
}
