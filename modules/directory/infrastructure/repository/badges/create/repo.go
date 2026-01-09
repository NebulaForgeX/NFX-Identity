package create

import (
	"nfxid/modules/directory/domain/badges"

	"gorm.io/gorm"
)

// Handler 处理创建操作，实现 badges.Create 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Create Handler
func NewHandler(db *gorm.DB) badges.Create {
	return &Handler{db: db}
}
