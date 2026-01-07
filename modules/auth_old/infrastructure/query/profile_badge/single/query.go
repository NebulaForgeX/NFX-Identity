package single

import (
	profileBadgeDomain "nfxid/modules/auth/domain/profile_badge"

	"gorm.io/gorm"
)

// Handler 处理单个查询操作，实现 profileBadgeDomain.Single 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Single Handler
func NewHandler(db *gorm.DB) profileBadgeDomain.Single {
	return &Handler{db: db}
}
