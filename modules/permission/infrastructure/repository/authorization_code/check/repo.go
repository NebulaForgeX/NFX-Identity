package check

import (
	authorizationCodeDomain "nfxid/modules/permission/domain/authorization_code"

	"gorm.io/gorm"
)

// Handler 处理检查操作，实现 authorizationCodeDomain.Check 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Check Handler
func NewHandler(db *gorm.DB) authorizationCodeDomain.Check {
	return &Handler{db: db}
}
