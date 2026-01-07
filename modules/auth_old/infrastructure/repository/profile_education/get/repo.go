package get

import (
	education "nfxid/modules/auth/domain/profile_education"

	"gorm.io/gorm"
)

// Handler 处理获取数据操作，实现 education.Get 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Get Handler
func NewHandler(db *gorm.DB) education.Get {
	return &Handler{db: db}
}
