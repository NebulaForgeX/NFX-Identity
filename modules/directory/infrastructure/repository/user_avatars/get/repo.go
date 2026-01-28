package get

import (
	"nfxid/modules/directory/domain/user_avatars"

	"gorm.io/gorm"
)

// Handler 处理获取操作，实现 user_avatars.Get 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Get Handler
func NewHandler(db *gorm.DB) user_avatars.Get {
	return &Handler{db: db}
}
