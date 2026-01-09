package delete

import (
	"nfxid/modules/image/domain/images"

	"gorm.io/gorm"
)

// Handler 处理删除操作，实现 images.Delete 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Delete Handler
func NewHandler(db *gorm.DB) images.Delete {
	return &Handler{db: db}
}
