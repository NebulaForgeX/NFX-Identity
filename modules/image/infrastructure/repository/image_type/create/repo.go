package create

import (
	imageTypeDomain "nfxid/modules/image/domain/image_type"

	"gorm.io/gorm"
)

// Handler 处理创建操作，实现 imageTypeDomain.Create 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Create Handler
func NewHandler(db *gorm.DB) imageTypeDomain.Create {
	return &Handler{db: db}
}
