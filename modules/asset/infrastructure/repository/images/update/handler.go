package update

import (
	imagesDomain "nfxidentity/modules/asset/domain/images"

	"gorm.io/gorm"
)

type Handler struct{ db *gorm.DB }

func NewHandler(db *gorm.DB) imagesDomain.Update { return &Handler{db: db} }
