package update

import (
	"nfxid/modules/image/domain/image_tags"

	"gorm.io/gorm"
)

// Handler 处理更新操作，实现 image_tags.Update 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Update Handler
func NewHandler(db *gorm.DB) image_tags.Update {
	return &Handler{db: db}
}
