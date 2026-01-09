package get

import (
	"nfxid/modules/directory/domain/user_preferences"

	"gorm.io/gorm"
)

// Handler 处理获取数据操作，实现 user_preferences.Get 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Get Handler
func NewHandler(db *gorm.DB) user_preferences.Get {
	return &Handler{db: db}
}
