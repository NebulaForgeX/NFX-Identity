package count

import (
	profileDomain "nfxid/modules/auth/domain/profile"

	"gorm.io/gorm"
)

// Handler 处理计数操作，实现 profileDomain.Count 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Count Handler
func NewHandler(db *gorm.DB) profileDomain.Count {
	return &Handler{db: db}
}
