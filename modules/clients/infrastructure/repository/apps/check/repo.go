package check

import (
	"nfxid/modules/clients/domain/apps"

	"gorm.io/gorm"
)

// Handler 处理检查操作，实现 apps.Check 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Check Handler
func NewHandler(db *gorm.DB) apps.Check {
	return &Handler{db: db}
}
