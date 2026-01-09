package create

import (
	"nfxid/modules/directory/domain/user_preferences"

	"gorm.io/gorm"
)

// Handler 处理创建操作，实现 user_preferences.Create 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Create Handler
func NewHandler(db *gorm.DB) user_preferences.Create {
	return &Handler{db: db}
}
