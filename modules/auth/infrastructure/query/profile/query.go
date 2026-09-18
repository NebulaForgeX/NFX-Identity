package profile

import (
	"nfxidentity/modules/auth/infrastructure/query/profile/list"
	profileQuery "nfxidentity/modules/auth/query/profile"

	"gorm.io/gorm"
)

func NewQuery(db *gorm.DB) *profileQuery.Query {
	return &profileQuery.Query{
		AuthorityList: list.NewAuthorityHandler(db),
		ForgerList:    list.NewForgerHandler(db),
	}
}
