package update

import (
	"nfxid/modules/auth/domain/account_lockouts"

	"gorm.io/gorm"
)

// Handler 处理更新操作，实现 account_lockouts.Update 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Update Handler
func NewHandler(db *gorm.DB) account_lockouts.Update {
	return &Handler{db: db}
}
