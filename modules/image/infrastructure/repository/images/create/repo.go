package create

import (
	"nfxid/modules/image/domain/images"

	"gorm.io/gorm"
)

// Handler 处理创建操作，实现 images.Create 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Create Handler
func NewHandler(db *gorm.DB) images.Create {
	return &Handler{db: db}
}
