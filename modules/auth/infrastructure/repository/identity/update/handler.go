package update

import (
	identityDomain "nfxidentity/modules/auth/domain/identity"

	"gorm.io/gorm"
)

type Handler struct{ db *gorm.DB }

func NewHandler(db *gorm.DB) identityDomain.Update { return &Handler{db: db} }
