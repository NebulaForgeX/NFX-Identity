package profile_education

import (
	educationDomain "nfxid/modules/auth/domain/profile_education"

	"gorm.io/gorm"
)

// Handler 处理查询操作，实现 education.Query 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Education Query Handler
func NewHandler(db *gorm.DB) educationDomain.Query {
	return &Handler{db: db}
}
