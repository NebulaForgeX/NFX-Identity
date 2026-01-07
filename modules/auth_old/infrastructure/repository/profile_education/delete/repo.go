package delete

import (
	education "nfxid/modules/auth/domain/profile_education"

	"gorm.io/gorm"
)

// Handler 处理删除操作，实现 education.Delete 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Delete Handler
func NewHandler(db *gorm.DB) education.Delete {
	return &Handler{db: db}
}
