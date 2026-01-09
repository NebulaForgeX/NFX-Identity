package delete

import (
	"nfxid/modules/image/domain/image_tags"

	"gorm.io/gorm"
)

// Handler 处理删除操作，实现 image_tags.Delete 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Delete Handler
func NewHandler(db *gorm.DB) image_tags.Delete {
	return &Handler{db: db}
}
