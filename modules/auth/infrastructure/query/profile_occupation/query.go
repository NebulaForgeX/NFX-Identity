package profile_occupation

import (
	occupationDomain "nfxid/modules/auth/domain/profile_occupation"

	"gorm.io/gorm"
)

// Handler 处理查询操作，实现 occupation.Query 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Occupation Query Handler
func NewHandler(db *gorm.DB) occupationDomain.Query {
	return &Handler{db: db}
}
