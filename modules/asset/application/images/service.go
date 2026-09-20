package images

import (
	imagesDomain "nfxidentity/modules/asset/domain/images"
	"nfxidentity/modules/asset/infrastructure/objectstore"
	imgsQuery "nfxidentity/modules/asset/query/images"
)

type Service struct {
	imageRepo  *imagesDomain.Repo
	imageQuery *imgsQuery.Query
	storage    *objectstore.Store
}

func NewService(
	imageRepo *imagesDomain.Repo,
	imageQuery *imgsQuery.Query,
	storage *objectstore.Store,
) *Service {
	return &Service{
		imageRepo:  imageRepo,
		imageQuery: imageQuery,
		storage:    storage,
	}
}
