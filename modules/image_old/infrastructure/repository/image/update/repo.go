package update

import (
	imageDomain "nfxid/modules/image/domain/image"

	"gorm.io/gorm"
)

// Handler 处理更新操作，实现 imageDomain.Update 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Update Handler
func NewHandler(db *gorm.DB) imageDomain.Update {
	return &Handler{db: db}
}
