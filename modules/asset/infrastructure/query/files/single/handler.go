package single

import (
	fls "nfxidentity/modules/asset/query/files"

	"gorm.io/gorm"
)

type Handler struct{ db *gorm.DB }

func NewHandler(db *gorm.DB) fls.Single { return &Handler{db: db} }
