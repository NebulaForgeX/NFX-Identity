package delete

import (
	"nfxid/modules/auth/domain/account_lockouts"

	"gorm.io/gorm"
)

// Handler 处理删除操作，实现 account_lockouts.Delete 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Delete Handler
func NewHandler(db *gorm.DB) account_lockouts.Delete {
	return &Handler{db: db}
}
