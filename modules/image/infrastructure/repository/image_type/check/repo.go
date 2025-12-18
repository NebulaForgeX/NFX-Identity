package check

import (
	imageTypeDomain "nfxid/modules/image/domain/image_type"

	"gorm.io/gorm"
)

// Handler 处理检查操作，实现 imageTypeDomain.Check 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Check Handler
func NewHandler(db *gorm.DB) imageTypeDomain.Check {
	return &Handler{db: db}
}
