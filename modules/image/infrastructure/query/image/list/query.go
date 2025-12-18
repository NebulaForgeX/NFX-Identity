package list

import (
	imageDomain "nfxid/modules/image/domain/image"

	"gorm.io/gorm"
)

// Handler 处理列表查询操作，实现 imageDomain.List 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 List Handler
func NewHandler(db *gorm.DB) imageDomain.List {
	return &Handler{db: db}
}
