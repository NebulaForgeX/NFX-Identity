package delete

import (
	authorizationCodeDomain "nfxid/modules/permission/domain/authorization_code"

	"gorm.io/gorm"
)

// Handler 处理删除操作，实现 authorizationCodeDomain.Delete 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Delete Handler
func NewHandler(db *gorm.DB) authorizationCodeDomain.Delete {
	return &Handler{db: db}
}
