package create

import (
	"nfxid/modules/directory/domain/users"

	"gorm.io/gorm"
)

// Handler 处理创建操作，实现 users.Create 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Create Handler
func NewHandler(db *gorm.DB) users.Create {
	return &Handler{db: db}
}
