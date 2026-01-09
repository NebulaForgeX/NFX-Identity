package update

import (
	"nfxid/modules/image/domain/images"

	"gorm.io/gorm"
)

// Handler 处理更新操作，实现 images.Update 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Update Handler
func NewHandler(db *gorm.DB) images.Update {
	return &Handler{db: db}
}
