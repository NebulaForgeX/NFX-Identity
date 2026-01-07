package get

import (
	"nfxid/modules/auth/domain/user"

	"gorm.io/gorm"
)

// Handler 处理获取数据操作，实现 user.Get 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Get Handler
func NewHandler(db *gorm.DB) user.Get {
	return &Handler{db: db}
}
