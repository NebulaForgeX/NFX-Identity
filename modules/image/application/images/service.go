package images

import (
	"nfxid/constants"
	imageDomain "nfxid/modules/image/domain/images"
)

type Service struct {
	imageRepo       *imageDomain.Repo
	storageBasePath string
}

func NewService(
	storageBasePath string,
	imageRepo *imageDomain.Repo,
) *Service {
	if storageBasePath == "" {
		storageBasePath = constants.StorageBasePath
	}
	return &Service{
		imageRepo:       imageRepo,
		storageBasePath: storageBasePath,
	}
}
