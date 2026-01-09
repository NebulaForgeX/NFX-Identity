package delete

import (
	"nfxid/modules/image/domain/image_variants"

	"gorm.io/gorm"
)

// Handler 处理删除操作，实现 image_variants.Delete 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Delete Handler
func NewHandler(db *gorm.DB) image_variants.Delete {
	return &Handler{db: db}
}
