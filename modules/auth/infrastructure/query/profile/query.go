package profile

import (
	"nfxidentity/modules/auth/infrastructure/query/profile/authority"
	"nfxidentity/modules/auth/infrastructure/query/profile/forger"
	profileQuery "nfxidentity/modules/auth/query/profile"

	"gorm.io/gorm"
)

func NewQuery(db *gorm.DB) *profileQuery.Query {
	return &profileQuery.Query{Forger: forger.NewHandler(db), Authority: authority.NewHandler(db)}
}
