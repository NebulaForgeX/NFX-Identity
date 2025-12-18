package single

import (
	userDomain "nfxid/modules/auth/domain/user"

	"gorm.io/gorm"
)

// Handler 处理单个查询操作，实现 userDomain.Single 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Single Handler
func NewHandler(db *gorm.DB) userDomain.Single {
	return &Handler{db: db}
}
