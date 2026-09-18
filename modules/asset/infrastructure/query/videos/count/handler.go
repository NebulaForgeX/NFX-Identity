package count

import (
	vids "nfxidentity/modules/asset/query/videos"

	"gorm.io/gorm"
)

type Handler struct{ db *gorm.DB }

func NewHandler(db *gorm.DB) vids.Count { return &Handler{db: db} }
