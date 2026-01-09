package create

import (
	"nfxid/modules/image/domain/image_types"

	"gorm.io/gorm"
)

// Handler 处理创建操作，实现 image_types.Create 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Create Handler
func NewHandler(db *gorm.DB) image_types.Create {
	return &Handler{db: db}
}
