package update

import (
	"nfxid/modules/image/domain/image_types"

	"gorm.io/gorm"
)

// Handler 处理更新操作，实现 image_types.Update 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Update Handler
func NewHandler(db *gorm.DB) image_types.Update {
	return &Handler{db: db}
}
