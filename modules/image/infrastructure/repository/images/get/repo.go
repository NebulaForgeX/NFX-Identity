package get

import (
	"nfxid/modules/image/domain/images"

	"gorm.io/gorm"
)

// Handler 处理获取数据操作，实现 images.Get 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Get Handler
func NewHandler(db *gorm.DB) images.Get {
	return &Handler{db: db}
}
