package image

import (
	"context"
	"nebulaid/modules/image/application/image/queries"
	"nebulaid/modules/image/application/image/views"

	"github.com/google/uuid"
)

func (s *Service) GetImage(ctx context.Context, imageID uuid.UUID) (views.ImageView, error) {
	domainView, err := s.imageQuery.GetByID(ctx, imageID)
	if err != nil {
		return views.ImageView{}, err
	}
	return views.ImageViewMapper(domainView), nil
}

type GetImageListResult struct {
	Items []views.ImageView
	Total int64
}

func (s *Service) GetImageList(ctx context.Context, q queries.ImageListQuery) (GetImageListResult, error) {
	q.Normalize()
	domainViews, total, err := s.imageQuery.GetList(ctx, q)
	if err != nil {
		return GetImageListResult{}, err
	}
	items := make([]views.ImageView, len(domainViews))
	for i, v := range domainViews {
		items[i] = views.ImageViewMapper(v)
	}
	return GetImageListResult{
		Items: items,
		Total: total,
	}, nil
}
