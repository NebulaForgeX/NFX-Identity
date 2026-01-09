package get

import (
	"nfxid/modules/access/domain/scopes"

	"gorm.io/gorm"
)

// Handler 处理获取数据操作，实现 scopes.Get 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Get Handler
func NewHandler(db *gorm.DB) scopes.Get {
	return &Handler{db: db}
}
