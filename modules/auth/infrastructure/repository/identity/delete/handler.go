package delete

import (
	identityDomain "nfxidentity/modules/auth/domain/identity"

	"gorm.io/gorm"
)

type Handler struct{ db *gorm.DB }

func NewHandler(db *gorm.DB) identityDomain.Delete { return &Handler{db: db} }
