package list

import (
	rdbviews "nfxidentity/modules/auth/infrastructure/rdb/views"
	profileQuery "nfxidentity/modules/auth/query/profile"
	"nfxidentity/pkgs/query"

	"gorm.io/gorm"
)

var authoritySearchQueryConfig = &query.QueryConfig{
	SearchConfig: &query.SearchConfig{
		Fields: []string{
			rdbviews.ListAuthorityProfileItemCols.DisplayName,
		},
		Operator: "ILIKE",
		Logic:    "OR",
	},
}

type AuthorityHandler struct {
	db *gorm.DB
}

func NewAuthorityHandler(db *gorm.DB) profileQuery.AuthorityList {
	return &AuthorityHandler{db: db}
}

var forgerSearchQueryConfig = &query.QueryConfig{
	SearchConfig: &query.SearchConfig{
		Fields: []string{
			rdbviews.ListForgerProfileItemCols.DisplayName,
		},
		Operator: "ILIKE",
		Logic:    "OR",
	},
}

type ForgerHandler struct {
	db *gorm.DB
}

func NewForgerHandler(db *gorm.DB) profileQuery.ForgerList {
	return &ForgerHandler{db: db}
}
