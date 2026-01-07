package image_type

import (
	imageTypeDomain "nfxid/modules/image/domain/image_type"
)

type Service struct {
	imageTypeRepo  *imageTypeDomain.Repo
	imageTypeQuery *imageTypeDomain.Query
}

func NewService(
	imageTypeRepo *imageTypeDomain.Repo,
	imageTypeQuery *imageTypeDomain.Query,
) *Service {
	return &Service{
		imageTypeRepo:  imageTypeRepo,
		imageTypeQuery: imageTypeQuery,
	}
}
