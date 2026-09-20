package images

import (
	"context"

	imgsQuery "nfxidentity/modules/asset/query/images"

	"github.com/google/uuid"
)

type GetInput struct {
	ID uuid.UUID
}

func (s *Service) Get(ctx context.Context, in GetInput) (*imgsQuery.ImageVO, error) {
	return s.imageQuery.Single.ByID(ctx, in.ID)
}

func (s *Service) GetMany(ctx context.Context, ids []string) ([]imgsQuery.ImageVO, error) {
	out := make([]imgsQuery.ImageVO, 0, len(ids))
	for _, raw := range ids {
		id, err := uuid.Parse(raw)
		if err != nil {
			continue
		}
		vo, err := s.Get(ctx, GetInput{ID: id})
		if err != nil {
			continue
		}
		out = append(out, *vo)
	}
	return out, nil
}
