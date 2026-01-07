package get

import (
	authorizationCodeDomain "nfxid/modules/permission/domain/authorization_code"

	"gorm.io/gorm"
)

// Handler 处理获取操作，实现 authorizationCodeDomain.Get 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Get Handler
func NewHandler(db *gorm.DB) authorizationCodeDomain.Get {
	return &Handler{db: db}
}
