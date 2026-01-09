package create

import (
	"nfxid/modules/image/domain/image_variants"

	"gorm.io/gorm"
)

// Handler 处理创建操作，实现 image_variants.Create 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Create Handler
func NewHandler(db *gorm.DB) image_variants.Create {
	return &Handler{db: db}
}
