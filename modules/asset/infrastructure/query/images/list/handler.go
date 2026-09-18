package list

import (
	rdbviews "nfxidentity/modules/asset/infrastructure/rdb/views"
	imgs "nfxidentity/modules/asset/query/images"
	"nfxidentity/pkgs/query"

	"gorm.io/gorm"
)

var imageQueryConfig = &query.QueryConfig{
	SearchConfig: &query.SearchConfig{
		Fields: []string{
			rdbviews.ImagesActiveViewCols.FilePath,
			rdbviews.ImagesActiveViewCols.FileName,
			rdbviews.ImagesActiveViewCols.MimeType,
		},
		Operator: "ILIKE",
		Logic:    "OR",
	},
}

type Handler struct{ db *gorm.DB }

func NewHandler(db *gorm.DB) imgs.List { return &Handler{db: db} }
