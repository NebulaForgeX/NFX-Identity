package update

import (
	filesDomain "nfxidentity/modules/asset/domain/files"

	"gorm.io/gorm"
)

type Handler struct{ db *gorm.DB }

func NewHandler(db *gorm.DB) filesDomain.Update { return &Handler{db: db} }
