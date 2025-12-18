package delete

import (
	imageDomain "nfxid/modules/image/domain/image"

	"gorm.io/gorm"
)

// Handler 处理删除操作，实现 imageDomain.Delete 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Delete Handler
func NewHandler(db *gorm.DB) imageDomain.Delete {
	return &Handler{db: db}
}
