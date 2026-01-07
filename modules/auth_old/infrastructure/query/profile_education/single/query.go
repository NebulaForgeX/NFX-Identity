package single

import (
	educationDomain "nfxid/modules/auth/domain/profile_education"

	"gorm.io/gorm"
)

// Handler 处理单个查询操作，实现 educationDomain.Single 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Single Handler
func NewHandler(db *gorm.DB) educationDomain.Single {
	return &Handler{db: db}
}
