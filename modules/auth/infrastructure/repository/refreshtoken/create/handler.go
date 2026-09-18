package create

import (
	refreshtokenDomain "nfxidentity/modules/auth/domain/refreshtoken"

	"gorm.io/gorm"
)

type Handler struct{ db *gorm.DB }

func NewHandler(db *gorm.DB) refreshtokenDomain.Create { return &Handler{db: db} }
