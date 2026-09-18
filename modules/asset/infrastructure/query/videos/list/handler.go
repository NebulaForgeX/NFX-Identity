package list

import (
	rdbviews "nfxidentity/modules/asset/infrastructure/rdb/views"
	vids "nfxidentity/modules/asset/query/videos"
	"nfxidentity/pkgs/query"

	"gorm.io/gorm"
)

var videoQueryConfig = &query.QueryConfig{
	SearchConfig: &query.SearchConfig{
		Fields: []string{
			rdbviews.VideosActiveViewCols.FilePath,
			rdbviews.VideosActiveViewCols.FileName,
			rdbviews.VideosActiveViewCols.MimeType,
		},
		Operator: "ILIKE",
		Logic:    "OR",
	},
}

type Handler struct{ db *gorm.DB }

func NewHandler(db *gorm.DB) vids.List { return &Handler{db: db} }
