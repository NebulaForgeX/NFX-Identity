package delete

import (
	imageTypeDomain "nfxid/modules/image/domain/image_type"

	"gorm.io/gorm"
)

// Handler 处理删除操作，实现 imageTypeDomain.Delete 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Delete Handler
func NewHandler(db *gorm.DB) imageTypeDomain.Delete {
	return &Handler{db: db}
}
