package check

import (
	"nfxid/modules/clients/domain/client_credentials"

	"gorm.io/gorm"
)

// Handler 处理检查操作，实现 client_credentials.Check 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Check Handler
func NewHandler(db *gorm.DB) client_credentials.Check {
	return &Handler{db: db}
}
