package check

import (
	"nfxid/modules/clients/domain/ip_allowlist"

	"gorm.io/gorm"
)

// Handler 处理检查操作，实现 ip_allowlist.Check 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Check Handler
func NewHandler(db *gorm.DB) ip_allowlist.Check {
	return &Handler{db: db}
}
