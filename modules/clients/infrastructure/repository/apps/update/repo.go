package update

import (
	"nfxid/modules/clients/domain/apps"

	"gorm.io/gorm"
)

// Handler 处理更新操作，实现 apps.Update 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Update Handler
func NewHandler(db *gorm.DB) apps.Update {
	return &Handler{db: db}
}
