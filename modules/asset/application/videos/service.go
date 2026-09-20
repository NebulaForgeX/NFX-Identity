package videos

import (
	videosDomain "nfxidentity/modules/asset/domain/videos"
	"nfxidentity/modules/asset/infrastructure/objectstore"
	vidsQuery "nfxidentity/modules/asset/query/videos"
)

type Service struct {
	videoRepo  *videosDomain.Repo
	videoQuery *vidsQuery.Query
	storage    *objectstore.Store
}

func NewService(
	videoRepo *videosDomain.Repo,
	videoQuery *vidsQuery.Query,
	storage *objectstore.Store,
) *Service {
	return &Service{
		videoRepo:  videoRepo,
		videoQuery: videoQuery,
		storage:    storage,
	}
}
