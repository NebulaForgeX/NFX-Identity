package update

import (
	"nfxid/modules/access/domain/tenant_roles"
	"gorm.io/gorm"
)

type Handler struct{ db *gorm.DB }

func NewHandler(db *gorm.DB) tenant_roles.Update {
	return &Handler{db: db}
}
