package audios

import (
	"context"

	audsQuery "nfxidentity/modules/asset/query/audios"
	"nfxidentity/pkgs/query"

	"github.com/google/uuid"
)

type ListInput struct {
	AccountID uuid.UUID
}

func (s *Service) List(ctx context.Context, in ListInput) ([]audsQuery.AudioVO, error) {
	page, err := s.audioQuery.List.Generic(ctx, audsQuery.ListQuery{
		DomainPagination: query.DomainPagination{All: true},
		UploaderIDs:      []uuid.UUID{in.AccountID},
	})
	if err != nil {
		return nil, err
	}
	return page.Items, nil
}
