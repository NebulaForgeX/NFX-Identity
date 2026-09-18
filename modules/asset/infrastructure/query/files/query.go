package files

import (
	"nfxidentity/modules/asset/infrastructure/query/files/count"
	"nfxidentity/modules/asset/infrastructure/query/files/list"
	"nfxidentity/modules/asset/infrastructure/query/files/single"
	fls "nfxidentity/modules/asset/query/files"

	"gorm.io/gorm"
)

func NewQuery(db *gorm.DB) *fls.Query {
	return &fls.Query{
		Single: single.NewHandler(db),
		List:   list.NewHandler(db),
		Count:  count.NewHandler(db),
	}
}
