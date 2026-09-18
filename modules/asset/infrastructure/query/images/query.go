package images

import (
	"nfxidentity/modules/asset/infrastructure/query/images/count"
	"nfxidentity/modules/asset/infrastructure/query/images/list"
	"nfxidentity/modules/asset/infrastructure/query/images/single"
	imgs "nfxidentity/modules/asset/query/images"

	"gorm.io/gorm"
)

func NewQuery(db *gorm.DB) *imgs.Query {
	return &imgs.Query{
		Single: single.NewHandler(db),
		List:   list.NewHandler(db),
		Count:  count.NewHandler(db),
	}
}
