package list

import (
	profileDomain "nfxid/modules/auth/domain/profile"

	"gorm.io/gorm"
)

// Handler 处理列表查询操作，实现 profileDomain.List 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 List Handler
func NewHandler(db *gorm.DB) profileDomain.List {
	return &Handler{db: db}
}
