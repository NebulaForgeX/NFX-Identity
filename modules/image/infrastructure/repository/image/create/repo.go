package create

import (
	imageDomain "nfxid/modules/image/domain/image"

	"gorm.io/gorm"
)

// Handler 处理创建操作，实现 imageDomain.Create 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Create Handler
func NewHandler(db *gorm.DB) imageDomain.Create {
	return &Handler{db: db}
}
