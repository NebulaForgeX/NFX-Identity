package count

import (
	occupationDomain "nfxid/modules/auth/domain/profile_occupation"

	"gorm.io/gorm"
)

// Handler 处理计数操作，实现 occupationDomain.Count 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Count Handler
func NewHandler(db *gorm.DB) occupationDomain.Count {
	return &Handler{db: db}
}
