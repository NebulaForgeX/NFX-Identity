package single

import (
	imageTypeDomain "nfxid/modules/image/domain/image_type"

	"gorm.io/gorm"
)

// Handler 处理单个查询操作，实现 imageTypeDomain.Single 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Single Handler
func NewHandler(db *gorm.DB) imageTypeDomain.Single {
	return &Handler{db: db}
}
