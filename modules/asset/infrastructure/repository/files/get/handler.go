package get

import (
	filesDomain "nfxidentity/modules/asset/domain/files"

	"gorm.io/gorm"
)

type Handler struct{ db *gorm.DB }

func NewHandler(db *gorm.DB) filesDomain.Get { return &Handler{db: db} }
