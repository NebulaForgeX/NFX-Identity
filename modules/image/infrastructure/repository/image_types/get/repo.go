package get

import (
	"nfxid/modules/image/domain/image_types"

	"gorm.io/gorm"
)

// Handler 处理获取数据操作，实现 image_types.Get 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Get Handler
func NewHandler(db *gorm.DB) image_types.Get {
	return &Handler{db: db}
}
