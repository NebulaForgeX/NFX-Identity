package check

import (
	videosDomain "nfxidentity/modules/asset/domain/videos"

	"gorm.io/gorm"
)

type Handler struct{ db *gorm.DB }

func NewHandler(db *gorm.DB) videosDomain.Check { return &Handler{db: db} }
