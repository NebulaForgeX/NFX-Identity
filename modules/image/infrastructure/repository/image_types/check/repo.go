package check

import (
	"nfxid/modules/image/domain/image_types"

	"gorm.io/gorm"
)

// Handler 处理检查操作，实现 image_types.Check 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Check Handler
func NewHandler(db *gorm.DB) image_types.Check {
	return &Handler{db: db}
}
