package list

import (
	userDomain "nfxid/modules/auth/domain/user"
	"nfxid/pkgs/query"

	"gorm.io/gorm"
)

var userQueryConfig = &query.QueryConfig{
	SearchConfig: &query.SearchConfig{
		Fields:   []string{"username", "email", "phone"},
		Operator: "ILIKE",
		Logic:    "OR",
	},
}

// Handler 处理列表查询操作，实现 userDomain.List 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 List Handler
func NewHandler(db *gorm.DB) userDomain.List {
	return &Handler{db: db}
}
