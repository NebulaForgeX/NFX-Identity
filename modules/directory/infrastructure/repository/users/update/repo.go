package update

import (
	"nfxid/modules/directory/domain/users"

	"gorm.io/gorm"
)

// Handler 处理更新操作，实现 users.Update 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Update Handler
func NewHandler(db *gorm.DB) users.Update {
	return &Handler{db: db}
}
