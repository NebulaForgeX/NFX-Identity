package create

import (
	videosDomain "nfxidentity/modules/asset/domain/videos"

	"gorm.io/gorm"
)

type Handler struct{ db *gorm.DB }

func NewHandler(db *gorm.DB) videosDomain.Create { return &Handler{db: db} }
