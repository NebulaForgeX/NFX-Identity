package count

import (
	fls "nfxidentity/modules/asset/query/files"

	"gorm.io/gorm"
)

type Handler struct{ db *gorm.DB }

func NewHandler(db *gorm.DB) fls.Count { return &Handler{db: db} }
