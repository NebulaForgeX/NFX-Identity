package get

import (
	refreshtokenDomain "nfxidentity/modules/auth/domain/refreshtoken"

	"gorm.io/gorm"
)

type Handler struct{ db *gorm.DB }

func NewHandler(db *gorm.DB) refreshtokenDomain.Get { return &Handler{db: db} }
