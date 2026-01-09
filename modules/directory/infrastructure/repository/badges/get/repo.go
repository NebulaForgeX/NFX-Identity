package get

import (
	"nfxid/modules/directory/domain/badges"

	"gorm.io/gorm"
)

// Handler 处理获取数据操作，实现 badges.Get 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Get Handler
func NewHandler(db *gorm.DB) badges.Get {
	return &Handler{db: db}
}
