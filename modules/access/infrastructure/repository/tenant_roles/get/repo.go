package get

import (
	"nfxid/modules/access/domain/tenant_roles"
	"gorm.io/gorm"
)

type Handler struct{ db *gorm.DB }

func NewHandler(db *gorm.DB) tenant_roles.Get {
	return &Handler{db: db}
}
