package single

import (
	imgs "nfxidentity/modules/asset/query/images"

	"gorm.io/gorm"
)

type Handler struct{ db *gorm.DB }

func NewHandler(db *gorm.DB) imgs.Single { return &Handler{db: db} }
