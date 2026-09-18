package get

import (
	audiosDomain "nfxidentity/modules/asset/domain/audios"

	"gorm.io/gorm"
)

type Handler struct{ db *gorm.DB }

func NewHandler(db *gorm.DB) audiosDomain.Get { return &Handler{db: db} }
