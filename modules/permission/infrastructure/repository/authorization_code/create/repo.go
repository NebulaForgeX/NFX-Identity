package create

import (
	authorizationCodeDomain "nfxid/modules/permission/domain/authorization_code"

	"gorm.io/gorm"
)

// Handler 处理创建操作，实现 authorizationCodeDomain.Create 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Create Handler
func NewHandler(db *gorm.DB) authorizationCodeDomain.Create {
	return &Handler{db: db}
}
