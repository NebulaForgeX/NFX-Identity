package update

import (
	"nfxid/modules/directory/domain/badges"

	"gorm.io/gorm"
)

// Handler 处理更新操作，实现 badges.Update 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Update Handler
func NewHandler(db *gorm.DB) badges.Update {
	return &Handler{db: db}
}
