package role

import (
	roleDomain "nfxid/modules/auth/domain/role"

	"gorm.io/gorm"
)

// Handler 处理查询操作，实现 role.Query 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Role Query Handler
func NewHandler(db *gorm.DB) roleDomain.Query {
	return &Handler{db: db}
}
