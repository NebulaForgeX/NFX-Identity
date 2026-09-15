package update

import (
	"nfxidentity/modules/asset/domain/images"

	"gorm.io/gorm"
)

type Handler struct{ db *gorm.DB }

func NewHandler(db *gorm.DB) images.Update { return &Handler{db: db} }
