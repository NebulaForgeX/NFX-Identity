package create

import (
	education "nfxid/modules/auth/domain/profile_education"

	"gorm.io/gorm"
)

// Handler 处理创建操作，实现 education.Create 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Create Handler
func NewHandler(db *gorm.DB) education.Create {
	return &Handler{db: db}
}
