package list

import (
	badgeDomain "nfxid/modules/auth/domain/badge"
	"nfxid/pkgs/query"

	"gorm.io/gorm"
)

var badgeQueryConfig = &query.QueryConfig{
	SearchConfig: &query.SearchConfig{
		Fields:   []string{"name", "description"},
		Operator: "ILIKE",
		Logic:    "OR",
	},
}

// Handler 处理列表查询操作，实现 badgeDomain.List 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 List Handler
func NewHandler(db *gorm.DB) badgeDomain.List {
	return &Handler{db: db}
}
