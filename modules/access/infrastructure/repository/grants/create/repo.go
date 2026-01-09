package create

import (
	"nfxid/modules/access/domain/grants"

	"gorm.io/gorm"
)

// Handler 处理创建操作，实现 grants.Create 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Create Handler
func NewHandler(db *gorm.DB) grants.Create {
	return &Handler{db: db}
}
