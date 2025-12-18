package single

import (
	profileDomain "nfxid/modules/auth/domain/profile"

	"gorm.io/gorm"
)

// Handler 处理单个查询操作，实现 profileDomain.Single 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Single Handler
func NewHandler(db *gorm.DB) profileDomain.Single {
	return &Handler{db: db}
}
