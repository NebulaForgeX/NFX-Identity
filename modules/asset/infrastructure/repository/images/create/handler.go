package create

import (
	imagesDomain "nfxidentity/modules/asset/domain/images"

	"gorm.io/gorm"
)

type Handler struct{ db *gorm.DB }

func NewHandler(db *gorm.DB) imagesDomain.Create { return &Handler{db: db} }
