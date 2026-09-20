package videos

import (
	"context"

	vidsQuery "nfxidentity/modules/asset/query/videos"
	"nfxidentity/pkgs/query"

	"github.com/google/uuid"
)

type ListInput struct {
	AccountID uuid.UUID
}

func (s *Service) List(ctx context.Context, in ListInput) ([]vidsQuery.VideoVO, error) {
	page, err := s.videoQuery.List.Generic(ctx, vidsQuery.ListQuery{
		DomainPagination: query.DomainPagination{All: true},
		UploaderIDs:      []uuid.UUID{in.AccountID},
	})
	if err != nil {
		return nil, err
	}
	return page.Items, nil
}
