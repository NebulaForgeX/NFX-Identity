package delete

import (
	filesDomain "nfxidentity/modules/asset/domain/files"

	"gorm.io/gorm"
)

type Handler struct{ db *gorm.DB }

func NewHandler(db *gorm.DB) filesDomain.Delete { return &Handler{db: db} }
