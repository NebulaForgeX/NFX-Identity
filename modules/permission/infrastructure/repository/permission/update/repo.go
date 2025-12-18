package update

import (
	permissionDomain "nfxid/modules/permission/domain/permission"

	"gorm.io/gorm"
)

// Handler 处理更新操作，实现 permissionDomain.Update 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Update Handler
func NewHandler(db *gorm.DB) permissionDomain.Update {
	return &Handler{db: db}
}
