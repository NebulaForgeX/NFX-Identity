package check

import (
	"nfxid/modules/access/domain/scopes"

	"gorm.io/gorm"
)

// Handler 处理检查操作，实现 scopes.Check 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Check Handler
func NewHandler(db *gorm.DB) scopes.Check {
	return &Handler{db: db}
}
