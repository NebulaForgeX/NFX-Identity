package user

import (
	userDomain "nfxid/modules/auth/domain/user"

	"gorm.io/gorm"
)

// Handler 处理查询操作，实现 user.Query 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 User Query Handler
func NewHandler(db *gorm.DB) userDomain.Query {
	return &Handler{db: db}
}
