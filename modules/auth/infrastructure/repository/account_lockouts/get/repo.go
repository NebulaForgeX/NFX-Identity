package get

import (
	"nfxid/modules/auth/domain/account_lockouts"

	"gorm.io/gorm"
)

// Handler 处理获取数据操作，实现 account_lockouts.Get 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Get Handler
func NewHandler(db *gorm.DB) account_lockouts.Get {
	return &Handler{db: db}
}
