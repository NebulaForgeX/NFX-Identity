package list

import (
	rdbviews "nfxidentity/modules/asset/infrastructure/rdb/views"
	fls "nfxidentity/modules/asset/query/files"
	"nfxidentity/pkgs/query"

	"gorm.io/gorm"
)

var fileQueryConfig = &query.QueryConfig{
	SearchConfig: &query.SearchConfig{
		Fields: []string{
			rdbviews.FilesActiveViewCols.FilePath,
			rdbviews.FilesActiveViewCols.FileName,
			rdbviews.FilesActiveViewCols.MimeType,
		},
		Operator: "ILIKE",
		Logic:    "OR",
	},
}

type Handler struct{ db *gorm.DB }

func NewHandler(db *gorm.DB) fls.List { return &Handler{db: db} }
