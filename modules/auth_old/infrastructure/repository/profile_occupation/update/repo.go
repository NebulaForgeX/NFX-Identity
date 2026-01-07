package update

import (
	occupation "nfxid/modules/auth/domain/profile_occupation"

	"gorm.io/gorm"
)

// Handler 处理更新操作，实现 occupation.Update 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Update Handler
func NewHandler(db *gorm.DB) occupation.Update {
	return &Handler{db: db}
}
