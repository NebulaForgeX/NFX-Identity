package delete

import (
	"nfxid/modules/directory/domain/user_preferences"

	"gorm.io/gorm"
)

// Handler 处理删除操作，实现 user_preferences.Delete 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Delete Handler
func NewHandler(db *gorm.DB) user_preferences.Delete {
	return &Handler{db: db}
}
