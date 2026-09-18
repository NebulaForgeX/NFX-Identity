package count

import (
	auds "nfxidentity/modules/asset/query/audios"

	"gorm.io/gorm"
)

type Handler struct{ db *gorm.DB }

func NewHandler(db *gorm.DB) auds.Count { return &Handler{db: db} }
