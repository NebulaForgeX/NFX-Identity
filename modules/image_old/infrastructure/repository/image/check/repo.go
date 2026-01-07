package check

import (
	imageDomain "nfxid/modules/image/domain/image"

	"gorm.io/gorm"
)

// Handler 处理检查操作，实现 imageDomain.Check 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Check Handler
func NewHandler(db *gorm.DB) imageDomain.Check {
	return &Handler{db: db}
}
