package get

import (
	videosDomain "nfxidentity/modules/asset/domain/videos"

	"gorm.io/gorm"
)

type Handler struct{ db *gorm.DB }

func NewHandler(db *gorm.DB) videosDomain.Get { return &Handler{db: db} }
