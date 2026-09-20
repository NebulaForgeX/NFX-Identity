package images

import (
	"context"

	imgsQuery "nfxidentity/modules/asset/query/images"
	"nfxidentity/pkgs/query"

	"github.com/google/uuid"
)

type ListInput struct {
	AccountID uuid.UUID
}

func (s *Service) List(ctx context.Context, in ListInput) ([]imgsQuery.ImageVO, error) {
	page, err := s.imageQuery.List.Generic(ctx, imgsQuery.ListQuery{
		DomainPagination: query.DomainPagination{All: true},
		UploaderIDs:      []uuid.UUID{in.AccountID},
	})
	if err != nil {
		return nil, err
	}
	return page.Items, nil
}
