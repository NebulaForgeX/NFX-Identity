package files

import (
	"context"

	flsQuery "nfxidentity/modules/asset/query/files"
	"nfxidentity/pkgs/query"

	"github.com/google/uuid"
)

type ListInput struct {
	AccountID uuid.UUID
}

func (s *Service) List(ctx context.Context, in ListInput) ([]flsQuery.FileVO, error) {
	page, err := s.fileQuery.List.Generic(ctx, flsQuery.ListQuery{
		DomainPagination: query.DomainPagination{All: true},
		UploaderIDs:      []uuid.UUID{in.AccountID},
	})
	if err != nil {
		return nil, err
	}
	return page.Items, nil
}
