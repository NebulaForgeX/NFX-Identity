package get

import (
	identityDomain "nfxidentity/modules/auth/domain/identity"

	"gorm.io/gorm"
)

type Handler struct{ db *gorm.DB }

func NewHandler(db *gorm.DB) identityDomain.Get { return &Handler{db: db} }
