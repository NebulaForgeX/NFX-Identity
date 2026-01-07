package check

import (
	education "nfxid/modules/auth/domain/profile_education"

	"gorm.io/gorm"
)

// Handler 处理检查操作，实现 education.Check 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Check Handler
func NewHandler(db *gorm.DB) education.Check {
	return &Handler{db: db}
}
