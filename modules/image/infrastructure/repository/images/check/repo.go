package check

import (
	"nfxid/modules/image/domain/images"

	"gorm.io/gorm"
)

// Handler 处理检查操作，实现 images.Check 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Check Handler
func NewHandler(db *gorm.DB) images.Check {
	return &Handler{db: db}
}
