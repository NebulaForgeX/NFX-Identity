package videos

import (
	"nfxidentity/modules/asset/infrastructure/query/videos/count"
	"nfxidentity/modules/asset/infrastructure/query/videos/list"
	"nfxidentity/modules/asset/infrastructure/query/videos/single"
	vids "nfxidentity/modules/asset/query/videos"

	"gorm.io/gorm"
)

func NewQuery(db *gorm.DB) *vids.Query {
	return &vids.Query{
		Single: single.NewHandler(db),
		List:   list.NewHandler(db),
		Count:  count.NewHandler(db),
	}
}
