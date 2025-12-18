package create

import (
	permissionDomain "nfxid/modules/permission/domain/permission"

	"gorm.io/gorm"
)

// Handler 处理创建操作，实现 permissionDomain.Create 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Create Handler
func NewHandler(db *gorm.DB) permissionDomain.Create {
	return &Handler{db: db}
}
