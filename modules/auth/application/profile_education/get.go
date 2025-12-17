package profile_education

import (
	"context"
	educationViews "nfxid/modules/auth/application/profile_education/views"
	educationDomain "nfxid/modules/auth/domain/profile_education"

	"github.com/google/uuid"
)

func (s *Service) GetEducation(ctx context.Context, educationID uuid.UUID) (educationViews.EducationView, error) {
	domainView, err := s.educationQuery.ByID(ctx, educationID)
	if err != nil {
		return educationViews.EducationView{}, err
	}
	return educationViews.EducationViewMapper(domainView), nil
}

func (s *Service) GetEducationsByProfileID(ctx context.Context, profileID uuid.UUID) ([]educationViews.EducationView, error) {
	domainViews, err := s.educationQuery.ByProfileID(ctx, profileID)
	if err != nil {
		return nil, err
	}
	result := make([]educationViews.EducationView, len(domainViews))
	for i, v := range domainViews {
		result[i] = educationViews.EducationViewMapper(v)
	}
	return result, nil
}

type GetEducationListResult struct {
	Items []educationViews.EducationView
	Total int64
}

func (s *Service) GetEducationList(ctx context.Context, q educationDomain.ListQuery) (GetEducationListResult, error) {
	q.Normalize()
	domainViews, total, err := s.educationQuery.List(ctx, q)
	if err != nil {
		return GetEducationListResult{}, err
	}
	items := make([]educationViews.EducationView, len(domainViews))
	for i, v := range domainViews {
		items[i] = educationViews.EducationViewMapper(v)
	}
	return GetEducationListResult{
		Items: items,
		Total: total,
	}, nil
}
