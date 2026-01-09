package check

import (
	"nfxid/modules/directory/domain/users"

	"gorm.io/gorm"
)

// Handler 处理检查操作，实现 users.Check 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Check Handler
func NewHandler(db *gorm.DB) users.Check {
	return &Handler{db: db}
}
