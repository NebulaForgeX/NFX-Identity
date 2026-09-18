package update

import (
	profileDomain "nfxidentity/modules/auth/domain/profile"

	"gorm.io/gorm"
)

type Handler struct{ db *gorm.DB }

func NewHandler(db *gorm.DB) profileDomain.Update {
	return &Handler{db: db}
}
