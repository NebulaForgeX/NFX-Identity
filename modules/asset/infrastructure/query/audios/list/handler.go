package list

import (
	rdbviews "nfxidentity/modules/asset/infrastructure/rdb/views"
	auds "nfxidentity/modules/asset/query/audios"
	"nfxidentity/pkgs/query"

	"gorm.io/gorm"
)

var videoQueryConfig = &query.QueryConfig{
	SearchConfig: &query.SearchConfig{
		Fields: []string{
			rdbviews.AudiosActiveViewCols.FilePath,
			rdbviews.AudiosActiveViewCols.FileName,
			rdbviews.AudiosActiveViewCols.MimeType,
		},
		Operator: "ILIKE",
		Logic:    "OR",
	},
}

type Handler struct{ db *gorm.DB }

func NewHandler(db *gorm.DB) auds.List { return &Handler{db: db} }
