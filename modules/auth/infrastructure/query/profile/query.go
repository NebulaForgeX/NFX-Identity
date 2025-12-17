package profile

import (
	profileDomain "nfxid/modules/auth/domain/profile"

	"gorm.io/gorm"
)

// Handler 处理查询操作，实现 profile.Query 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Profile Query Handler
func NewHandler(db *gorm.DB) profileDomain.Query {
	return &Handler{db: db}
}
