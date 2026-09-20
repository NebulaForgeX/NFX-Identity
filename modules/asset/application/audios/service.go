package audios

import (
	audiosDomain "nfxidentity/modules/asset/domain/audios"
	"nfxidentity/modules/asset/infrastructure/objectstore"
	audsQuery "nfxidentity/modules/asset/query/audios"
)

type Service struct {
	audioRepo  *audiosDomain.Repo
	audioQuery *audsQuery.Query
	storage    *objectstore.Store
}

func NewService(
	audioRepo *audiosDomain.Repo,
	audioQuery *audsQuery.Query,
	storage *objectstore.Store,
) *Service {
	return &Service{
		audioRepo:  audioRepo,
		audioQuery: audioQuery,
		storage:    storage,
	}
}
