package get

import (
	imageTypeDomain "nfxid/modules/image/domain/image_type"

	"gorm.io/gorm"
)

// Handler 处理获取数据操作，实现 imageTypeDomain.Get 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Get Handler
func NewHandler(db *gorm.DB) imageTypeDomain.Get {
	return &Handler{db: db}
}
