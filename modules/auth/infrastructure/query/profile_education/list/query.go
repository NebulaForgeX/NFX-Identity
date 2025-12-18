package list

import (
	educationDomain "nfxid/modules/auth/domain/profile_education"

	"gorm.io/gorm"
)

// Handler 处理列表查询操作，实现 educationDomain.List 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 List Handler
func NewHandler(db *gorm.DB) educationDomain.List {
	return &Handler{db: db}
}
