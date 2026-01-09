package check

import (
	"nfxid/modules/auth/domain/trusted_devices"

	"gorm.io/gorm"
)

// Handler 处理检查操作，实现 trusted_devices.Check 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Check Handler
func NewHandler(db *gorm.DB) trusted_devices.Check {
	return &Handler{db: db}
}
