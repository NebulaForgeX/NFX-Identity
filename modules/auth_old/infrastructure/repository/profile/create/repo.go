package create

import (
	"nfxid/modules/auth/domain/profile"

	"gorm.io/gorm"
)

// Handler 处理创建操作，实现 profile.Create 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Create Handler
func NewHandler(db *gorm.DB) profile.Create {
	return &Handler{db: db}
}
