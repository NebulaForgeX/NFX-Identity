package check

import (
	"nfxid/modules/directory/domain/user_avatars"

	"gorm.io/gorm"
)

// Handler 处理检查操作，实现 user_avatars.Check 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Check Handler
func NewHandler(db *gorm.DB) user_avatars.Check {
	return &Handler{db: db}
}
