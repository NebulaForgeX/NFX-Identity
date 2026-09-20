package files

import (
	filesDomain "nfxidentity/modules/asset/domain/files"
	"nfxidentity/modules/asset/infrastructure/objectstore"
	flsQuery "nfxidentity/modules/asset/query/files"
)

type Service struct {
	fileRepo  *filesDomain.Repo
	fileQuery *flsQuery.Query
	storage   *objectstore.Store
}

func NewService(
	fileRepo *filesDomain.Repo,
	fileQuery *flsQuery.Query,
	storage *objectstore.Store,
) *Service {
	return &Service{
		fileRepo:  fileRepo,
		fileQuery: fileQuery,
		storage:   storage,
	}
}
