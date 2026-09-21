package single

import (
	accountQuery "nfxidentity/modules/auth/query/account"

	"gorm.io/gorm"
)

type Handler struct{ db *gorm.DB }

func NewHandler(db *gorm.DB) accountQuery.Single { return &Handler{db: db} }
