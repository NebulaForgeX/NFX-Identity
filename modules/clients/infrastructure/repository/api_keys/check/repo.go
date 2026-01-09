package check

import (
	"nfxid/modules/clients/domain/api_keys"

	"gorm.io/gorm"
)

// Handler 处理检查操作，实现 api_keys.Check 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Check Handler
func NewHandler(db *gorm.DB) api_keys.Check {
	return &Handler{db: db}
}
