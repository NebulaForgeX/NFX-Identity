package badge

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

// Handler 处理查询操作，实现 badge.Query 接口
type Handler struct {
	db *gorm.DB
}

// NewHandler 创建新的 Badge Query Handler
func NewHandler(db *gorm.DB) badgeDomain.Query {
	return &Handler{db: db}
}
