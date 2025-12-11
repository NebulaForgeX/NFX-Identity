package image_type

import (
	"context"
	"nfxid/modules/image/application/image_type/queries"
	"nfxid/modules/image/application/image_type/views"

	"github.com/google/uuid"
)

func (s *Service) GetImageType(ctx context.Context, imageTypeID uuid.UUID) (views.ImageTypeView, error) {
	domainView, err := s.imageTypeQuery.GetByID(ctx, imageTypeID)
	if err != nil {
		return views.ImageTypeView{}, err
	}
	return views.ImageTypeViewMapper(domainView), nil
}

func (s *Service) GetImageTypeByKey(ctx context.Context, key string) (views.ImageTypeView, error) {
	domainView, err := s.imageTypeQuery.GetByKey(ctx, key)
	if err != nil {
		return views.ImageTypeView{}, err
	}
	return views.ImageTypeViewMapper(domainView), nil
}

type GetImageTypeListResult struct {
	Items []views.ImageTypeView
	Total int64
}

func (s *Service) GetImageTypeList(ctx context.Context, q queries.ImageTypeListQuery) (GetImageTypeListResult, error) {
	q.Normalize()
	domainViews, total, err := s.imageTypeQuery.GetList(ctx, q)
	if err != nil {
		return GetImageTypeListResult{}, err
	}
	items := make([]views.ImageTypeView, len(domainViews))
	for i, v := range domainViews {
		items[i] = views.ImageTypeViewMapper(v)
	}
	return GetImageTypeListResult{
		Items: items,
		Total: total,
	}, nil
}
