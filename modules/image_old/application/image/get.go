package image

import (
	"context"
	"nfxid/modules/image/application/image/views"
	imageDomain "nfxid/modules/image/domain/image"

	"github.com/google/uuid"
)

func (s *Service) GetImage(ctx context.Context, imageID uuid.UUID) (views.ImageView, error) {
	domainView, err := s.imageQuery.Single.ByID(ctx, imageID)
	if err != nil {
		return views.ImageView{}, err
	}
	return views.ImageViewMapper(*domainView), nil
}

type GetImageListResult struct {
	Items []views.ImageView
	Total int64
}

func (s *Service) GetImageList(ctx context.Context, q imageDomain.ListQuery) (GetImageListResult, error) {
	q.Normalize()
	domainViews, total, err := s.imageQuery.List.Generic(ctx, q)
	if err != nil {
		return GetImageListResult{}, err
	}
	items := make([]views.ImageView, len(domainViews))
	for i, v := range domainViews {
		items[i] = views.ImageViewMapper(*v)
	}
	return GetImageListResult{
		Items: items,
		Total: total,
	}, nil
}
