package image_type

import (
	imageTypeQueries "nfxid/modules/image/application/image_type/queries"
	imageTypeDomain "nfxid/modules/image/domain/image_type"
)

type Service struct {
	imageTypeRepo  imageTypeDomain.Repo
	imageTypeQuery imageTypeQueries.ImageTypeQuery
}

func NewService(
	imageTypeRepo imageTypeDomain.Repo,
	imageTypeQuery imageTypeQueries.ImageTypeQuery,
) *Service {
	return &Service{
		imageTypeRepo:  imageTypeRepo,
		imageTypeQuery: imageTypeQuery,
	}
}
