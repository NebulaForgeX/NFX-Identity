package image_type

import (
	imageTypeDomain "nfxid/modules/image/domain/image_type"
	"nfxid/modules/image/infrastructure/query/image_type/list"
	"nfxid/modules/image/infrastructure/query/image_type/single"

	"gorm.io/gorm"
)

// NewHandler 创建新的 ImageType Query Handler
func NewHandler(db *gorm.DB) *imageTypeDomain.Query {
	return &imageTypeDomain.Query{
		Single: single.NewHandler(db),
		List:   list.NewHandler(db),
	}
}
