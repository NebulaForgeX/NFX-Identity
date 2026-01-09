package get

import (
	"nfxid/modules/access/domain/grants"

	"gorm.io/gorm"
)

// Handler 处理获取数据操作，实现 grants.Get 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Get Handler
func NewHandler(db *gorm.DB) grants.Get {
	return &Handler{db: db}
}
